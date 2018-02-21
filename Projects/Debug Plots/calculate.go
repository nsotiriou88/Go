/*
Package calculator unmarhals data stream from the UDP stream calculates metrics and emits them as models.Vectors.
*/
package calculator

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"sort"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/sportabletech/go-libs/uwb/packet"
	"github.com/sportabletech/go-libs/uwb/tlv"
	"github.com/sportabletech/uwb-metrics-server/models"
)

var timeChecked = false
var timeCorrection float64

// Demultiplex processes the (json) UDP stream. It Determines the ID of each input packet, in order to access its settings, and pass these to the calculators.
//
// Emit the results of those calculations on each output channel.
// Tag calibration is a pointer and may be modified by other goroutines, eg. http request.
// Always run in a seperate goroutine.
func Demultiplex(input chan []byte, outputs []chan models.Vectors, tagCalibration *models.Devices, appCfg *models.AppConfig, sessionCfg *models.Sessions) {
	var workers = make(map[string]chan packet.Packet)
	var loggers = make(map[string]chan packet.Packet)
	log.Println("Demultiplex (json) routine started")

	//combined tag data text logs
	logBuffer := make(chan packet.Packet, 10000)
	if appCfg.CombinedLogs.Enabled {
		go LogTagData(logBuffer, "combined", appCfg.CombinedLogs.Distances)
	}

	for data := range input {

		//	color.Yellow(string(data))

		scanner := bufio.NewScanner(bytes.NewBuffer(data)) //this allows us to buffer by line to extract multiple openrtls strings;
		for scanner.Scan() {
			var p packet.Packet
			line := scanner.Bytes()
			err := json.Unmarshal(line, &p)
			if err != nil {
				log.Println(err)
				//	log.Println("line ", string(line))
				continue //just carry on if error with Unmarshal
			}

			//check if first message from tag, if yes, spawn DoWork routine
			destination, exists := workers[p.ID]
			if !exists {
				destination = make(chan packet.Packet)
				workers[p.ID] = destination
				go DoWork(destination, outputs, tagCalibration, appCfg)
			}

			destination <- p

			//if logs for individual tags enabled & there is an open session from the webUI.
			if appCfg.TagLogs.Enabled && sessionCfg.PlayActive {
				tagLog, loggerExists := loggers[p.ID]
				if !loggerExists { //spawn new logger goroutine if new tag
					tagLog = make(chan packet.Packet)
					loggers[p.ID] = tagLog
					go LogTagData(tagLog, p.ID, appCfg.TagLogs.Distances)

				}

				tagLog <- p
				//	if sessionCfg.SessionActive {
				//		close(tagLog)
				//	}
			}

			if appCfg.CombinedLogs.Enabled {
				if appCfg.CombinedLogs.UserAndIMUonly { //skip log if only userdata / imu wanted and not present
					if p.UserData == nil && p.Sensors == nil {
						continue
					}
				}
				logBuffer <- p
			}
		}
	}
}

//DemultiplexTLV is a development function - it will be merged back into the main Demultiplex function.
func DemultiplexTLV(input chan []byte, outputs []chan models.Vectors, tagCalibration *models.Devices, appCfg *models.AppConfig, sessionCfg *models.Sessions) {
	//	binDumpFile, err := os.Create("devLogs/TLVdump.dat")
	//	check(err)
	//	defer binDumpFile.Close()
	//	textDumpFile, err := os.Create("devLogs/TLVdump.txt")
	//	check(err)
	//	defer textDumpFile.Close()

	//tlv.Test()

	var workers = make(map[string]chan packet.Packet) //map of workers to spawn DoWork for each tag
	var loggers = make(map[string]chan packet.Packet)
	log.Println("Demultiplex (tlv) routine started")

	//combined tag data text logs
	logBuffer := make(chan packet.Packet, 10000)
	if appCfg.CombinedLogs.Enabled {
		go LogTagData(logBuffer, "combined", appCfg.CombinedLogs.Distances)
	}

	for data := range input {

		//	//THIS IS DEV DEBUG
		//	binDumpFile.Write(data)
		//	hexdata := fmt.Sprintf("\n%X", data)
		//	textDumpFile.Write([]byte(hexdata))
		//	//fmt.Printf(hexdata)
		//	//END DEV DEBUG

		var p []packet.Packet
		//fmt.Printf("\n%X", string(data))
		//if using binary (tlv) protocol
		err := tlv.UnmarshalTLV(data, &p)
		if err != nil {
			fmt.Printf("\n%X", string(data))
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		//end protocol fomrat if

		for _, pi := range p {

			destination, exists := workers[pi.ID]
			if !exists { //spawn new worker goroutine if new tag
				destination = make(chan packet.Packet)
				workers[pi.ID] = destination
				go DoWork(destination, outputs, tagCalibration, appCfg)
			}
			destination <- pi

			//if logs for individual tags enabled & there is an open session from the webUI.
			if appCfg.TagLogs.Enabled && sessionCfg.SessionActive {
				tagLog, loggerExists := loggers[pi.ID]
				if !loggerExists { //spawn new logger goroutine if new tag
					tagLog = make(chan packet.Packet)
					loggers[pi.ID] = tagLog
					go LogTagData(tagLog, pi.ID, appCfg.TagLogs.Distances)
				}

				tagLog <- pi
				//	if sessionCfg.SessionActive {
				//		close(tagLog)
				//	}
			}

			if appCfg.CombinedLogs.Enabled {
				logBuffer <- pi
			}
		}
	}
	log.Println("[!] Demultiplex (tlv) routine ENDED!!")
}

type measDiags struct {
	Samples             uint
	CalcCount           uint //used to slow down printouts
	VmaxExcededCount    float64
	OriginPosCount      float64
	TimeDeltaSmallCount float64
	ValidCount          float64
	MeasurementSum      float64
	ValidPercent        float64
	OutOfOrderCount     uint //used to count successive out of order packets (detects restart of master/ASE)
}

//Reset clears all diags, excluding ValidPercent
func (diags *measDiags) Reset() {
	diags.Samples = 0
	diags.VmaxExcededCount = 0
	diags.OriginPosCount = 0
	diags.TimeDeltaSmallCount = 0
	diags.ValidCount = 0
	diags.MeasurementSum = 0

}

//ResetHard clears all diags, including ValidPercent
func (diags *measDiags) ResetHard() {
	diags.Samples = 0
	diags.VmaxExcededCount = 0
	diags.OriginPosCount = 0
	diags.TimeDeltaSmallCount = 0
	diags.ValidCount = 0
	diags.MeasurementSum = 0
	diags.ValidPercent = 0
}

// DoWork is run as goroutine dedicated to do the calculations for a specific tag ID.
func DoWork(input chan packet.Packet, outputs []chan models.Vectors, settings *models.Devices, appCfg *models.AppConfig) {

	// Creating the map for the anchors and initialisations.
	count := 0
	plotData := make(map[string]map[string][]float64)
	plotData["dist"] = make(map[string][]float64)
	plotData["rssi"] = make(map[string][]float64)
	plotData["rssi2"] = make(map[string][]float64)

	//inputChannelIterator:
	for p := range input {

		// log.Println(string(line))
		//	log.Print(p.ID, p.Coordinates)
		for i := range p.Meas {
			// fmt.Println("a:", p.Meas[i].Anchor, "dist:", p.Meas[i].Dist, "rssi:", p.Meas[i].Rssi, "dB")
			// Register data to Map for plotting.
			plotData["dist"][p.Meas[i].Anchor] = append(plotData["dist"][p.Meas[i].Anchor], float64(p.Meas[i].Dist))
			plotData["rssi"][p.Meas[i].Anchor] = append(plotData["rssi"][p.Meas[i].Anchor], float64(p.Meas[i].Rssi))
			plotData["rssi2"][p.Meas[i].Anchor] = append(plotData["rssi2"][p.Meas[i].Anchor], float64(p.Meas[i].Rssi2))

		}

		// Print every 100 samples (5 secs)
		count++
		if count%100 == 99 {
			go Plot(p.ID, plotData)
		}

		if !timeChecked {
			timeCorrection = warnIfTimstampOld(p.Timestamp)
			timeChecked = true //we only check the first sample.

			if appCfg.OverRideTimeStamp {
				log.Printf("[!] correcting incoming timestamps by %3.4f (based on first packet seconds lag)", timeCorrection)
			}
		}

		outVector := models.NewVectors(p)

		//generate output channel
		for _, output := range outputs {
			output <- outVector
			//log.Println("pos valid", outVector.PositionValid)
			//	log.Println("[?]", soft.Timestamp, soft.Pos)
		}

		//	log.Printf("lag %3.3f", outVector.TsLocal-outVector.Timestamp)

	} //end inputChannelIterator loop.
}

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func warnIfTimstampOld(time float64) (correction float64) {
	correction = 0
	ageInSeconds := models.Timestamp() - time
	if ageInSeconds > 0.5 {
		log.Printf("[!] First Timestamp INVALID!! %3.4f seconds lag", ageInSeconds)
		log.Println("IS NTP SERVER CONFIGURED? DOES MASTER KNOW ABOUT IT!")
		correction = ageInSeconds
	} else {
		log.Printf("[+] timestamp %3.4f seconds lag", ageInSeconds)
	}

	return correction
}

// LogTagData writes each packet received as a json line of a log file.
func LogTagData(input chan packet.Packet, description string, bLogDistances bool) {
	const maxLogSize int64 = 50 * 1000 * 1000 //size in bytes (limit to 50-100MB for ease of use.)
	var file io.WriteCloser
	var err error
	var empty []packet.Meas
	parts := 0
	dirName := fmt.Sprintf("logs/%v/", time.Now().Format("20060102"))
	filePathName := dirName + fmt.Sprintf("%v_%v_p%v.txt", description, time.Now().Format("150405"), parts)
	log.Println("[+] created log file", filePathName)

	//check if dir exists, else create
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.Mkdir(dirName, 0777)
		log.Println("[+] created log directory", dirName)
	}

	//create file
	file, err = os.Create(filePathName)

	encoder := json.NewEncoder(file)
	check(err)

	for message := range input {

		if !bLogDistances {
			message.Meas = empty
		}

		err := encoder.Encode(message)
		if err != nil {
			log.Println(err)
			fmt.Println("LogTagData")
			fmt.Println(message)
		}
		fileInfo, err := os.Stat(filePathName)
		check(err)

		//rotate when too big
		if fileInfo.Size() > maxLogSize {
			file.Close()
			parts++
			//nb! if a session spans past midnight the logs will be in the same folder (i.e. date of start of session)
			filePathName = dirName + fmt.Sprintf("%v_%v_p%v.txt", description, time.Now().Format("150405"), parts)
			file, err = os.Create(filePathName)
			check(err)
			encoder = json.NewEncoder(file)
		}
	}
}

//Plot produces some useful diagnostic plots
func Plot(tagID string, data map[string]map[string][]float64) {

	path := "plots/" + tagID + "/"

	//remove directory and contents, recreate and create plots, index.html
	err := os.RemoveAll(path)
	if err != nil {
		log.Println("[!] failed to remove plot directory", path)
	} else {
		os.Mkdir(path, 0777)
		log.Println("[+] created plot directory", path)
	}

	var anchors []string

	// Draw all 3 sets of values for the histogram
	for measKeys, mapOfAnchors := range data { // Range over Dist, Rssi & Rssi2
		for anchor := range mapOfAnchors { // Range over individual Anchors

			if measKeys == "dist" {
				anchors = append(anchors, anchor)
			}
			v := plotter.Values(data[measKeys][anchor])

			// Make a plot and set its title.
			p, err := plot.New()
			if err != nil {
				panic(err)
			}

			// Create a histogram of our values
			h, err := plotter.NewHist(v, 50)
			if err != nil {
				panic(err)
			}

			// Normalize the area under the histogram to sum to one.
			h.Normalize(1)
			p.Add(h)

			// The normal distribution function (uncommend if you wish to display):
			// norm := plotter.NewFunction(distuv.UnitNormal.Prob)
			// norm.Color = color.RGBA{R: 255, A: 255}
			// norm.Width = vg.Points(2)
			// p.Add(norm)

			// Save the plots to PNG files.
			p.Title.Text = measKeys + " for anchor: " + anchor
			if err := p.Save(2.5*vg.Inch, 2.5*vg.Inch, path+measKeys+"_"+anchor+".png"); err != nil {
				log.Println("[!] error creating plot", err)
			}

		}
	}
	sort.Strings(anchors)
	createIndexPage(path, tagID, anchors)

}

func createIndexPage(path, tagID string, anchors []string) {

	fmt.Println("tag:", tagID, "anchors:", anchors)

	t, err := template.ParseFiles("indexTemplate.html")
	if err != nil {
		panic(err)
	}
	timeStr := "Tag " + tagID + " Debug Plots  -  " + (time.Now().Format(time.RFC850))
	titleStr := tagID + " plots "
	descrStr := "Distanses, RSSI and RSSI2 from all visible Anchors"
	content := indexContent{Title: titleStr, Header: timeStr, Description: descrStr, Anchors: anchors}

	//create file
	fo, err := os.Create(path + "index.html")
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	//err = t.ExecuteTemplate(os.Stdout, "index.html", content)
	err = t.ExecuteTemplate(fo, "indexTemplate.html", content)
	if err != nil {
		panic(err)
	}
}

type indexContent struct {
	Title       string
	Header      string
	Description string
	Anchors     []string
}
