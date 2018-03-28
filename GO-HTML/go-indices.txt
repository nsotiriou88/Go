// DoWork is run as goroutine dedicated to do the calculations for a specific tag ID.
func DoWork(input chan packet.Packet, outputs []chan models.Vectors, settings *models.Devices, appCfg *models.AppConfig, reset *bool) {

	// Creating the map for the anchors and initialisations.
	count := 0
	plotData := make(map[string]map[string][]float64)
	plotData["dist"] = make(map[string][]float64)
	plotData["rssi"] = make(map[string][]float64)
	plotData["rssi2"] = make(map[string][]float64)
	plotData["tqf"] = make(map[string][]float64)
	plotData["tqf2"] = make(map[string][]float64)

	//inputChannelIterator:
	for p := range input {

		// log.Println(string(line))
		//	log.Print(p.ID, p.Coordinates)
		for i := range p.Meas {
			// fmt.Println("a:", p.Meas[i].Anchor, "dist:", p.Meas[i].Dist, "rssi:", "dB", p.Meas[i].Rssi, "dB")
			// fmt.Println("tqf:", p.Meas[i].Tqf, "tqf2:", p.Meas[i].Tqf2)
			// Register data to Map for plotting.
			plotData["dist"][p.Meas[i].Anchor] = append(plotData["dist"][p.Meas[i].Anchor], float64(p.Meas[i].Dist))
			plotData["rssi"][p.Meas[i].Anchor] = append(plotData["rssi"][p.Meas[i].Anchor], float64(p.Meas[i].Rssi))
			plotData["rssi2"][p.Meas[i].Anchor] = append(plotData["rssi2"][p.Meas[i].Anchor], float64(p.Meas[i].Rssi2))
			plotData["tqf"][p.Meas[i].Anchor] = append(plotData["tqf"][p.Meas[i].Anchor], float64(p.Meas[i].Tqf))
			plotData["tqf2"][p.Meas[i].Anchor] = append(plotData["tqf2"][p.Meas[i].Anchor], float64(p.Meas[i].Tqf2))

		}

		// Print every 100 samples (5 secs) and clear map
		count++
		if count%100 == 0 {
			Plot(p.ID, plotData, reset)
			if *reset {
				plotData["dist"] = make(map[string][]float64)
				plotData["rssi"] = make(map[string][]float64)
				plotData["rssi2"] = make(map[string][]float64)
				plotData["tqf"] = make(map[string][]float64)
				plotData["tqf2"] = make(map[string][]float64)
			}
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
func Plot(tagID string, data map[string]map[string][]float64, reset *bool) {

	// Check if tagID is not registered
	flagLive := false
	for _, val := range LiveTags {
		if val == tagID {
			flagLive = true
			break
		}
	}

	if flagLive != true {
		LiveTags = append(LiveTags, tagID)
	}

	// Create the saving path
	os.Mkdir("plots/", 0777)
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
	samples := make(map[string]string)

	// Draw all 5 sets of values for the histogram
	for measKeys, mapOfAnchors := range data { // Range over Dist, Rssi, Rssi2, TQF & TQF2
		for anchor := range mapOfAnchors { // Range over individual Anchors

			// Create the list of anchors and calculate median & number of samples
			var meanStr, medianStr string
			if measKeys == "dist" {
				anchors = append(anchors, anchor)
				median := Median(data[measKeys][anchor])
				medianStr = strconv.FormatFloat(median, 'g', 4, 64)
				samples[anchor] = strconv.Itoa(len(data["dist"][anchor]))
				// fmt.Printf("It is type: %T \n", samples[anchor])
			}
			// Calculate mean for all measurements
			mean := stat.Mean(data[measKeys][anchor], nil)
			meanStr = strconv.FormatFloat(mean, 'g', 4, 64)

			// Plotting for dist, rssi and rssi2
			if measKeys == "dist" || measKeys == "rssi" || measKeys == "rssi2" {
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

				// Uncomment to Normalize the histogram
				// h.Normalize(1)
				p.Add(h)

				// The normal distribution function (uncommend if you wish to display):
				// norm := plotter.NewFunction(distuv.UnitNormal.Prob)
				// norm.Color = color.RGBA{R: 255, A: 255}
				// norm.Width = vg.Points(2)
				// p.Add(norm)

				// Save the plots to PNG files.
				if measKeys == "dist" {
					p.Title.Text = measKeys + " in m for anchor: " + anchor + "\n Mean: " + meanStr + "m, Median: " + medianStr + "m"
					if err := p.Save(2.5*vg.Inch, 2.5*vg.Inch, path+measKeys+"_"+anchor+".png"); err != nil {
						log.Println("[!] error creating plot", err)
					}
				} else {
					p.Title.Text = measKeys + " in db for anchor: " + anchor + "\n Mean: " + meanStr + "db"
					if err := p.Save(2.5*vg.Inch, 2.5*vg.Inch, path+measKeys+"_"+anchor+".png"); err != nil {
						log.Println("[!] error creating plot", err)
					}
				}

			} else { //Plotting for tqf and tqf2
				var group []float64
				var count1s, count2s int
				// Create a group with the frequency of 1s and 2s
				for _, freq := range data[measKeys][anchor] {
					if freq == 1 {
						count1s++
					} else {
						count2s++
					}
				}
				group = append(group, float64(count1s))
				group = append(group, float64(count2s))

				// Make a plot and set its title.
				v := plotter.Values(group)
				p, err := plot.New()
				if err != nil {
					panic(err)
				}
				p.Y.Label.Text = "Frequency"
				w := vg.Points(50)

				bars, err := plotter.NewBarChart(v, w)
				if err != nil {
					panic(err)
				}
				bars.LineStyle.Width = vg.Length(0)
				bars.Color = plotutil.Color(2)

				p.Add(bars)
				p.Legend.Add("Freq "+measKeys, bars)
				p.Legend.Top = true
				p.NominalX("1", "2")

				p.Title.Text = measKeys + " for anchor: " + anchor + "\n Mean: " + meanStr
				if err := p.Save(2.5*vg.Inch, 2.5*vg.Inch, path+measKeys+"_"+anchor+".png"); err != nil {
					log.Println("[!] error creating plot", err)
				}
			}

		}
	}

	sort.Strings(anchors)
	createIndexPage(path, tagID, anchors, samples, reset)
	createTagsIndexPage(LiveTags)

}

func createIndexPage(path, tagID string, anchors []string, samples map[string]string, reset *bool) {

	fmt.Println("tag:", tagID, "anchors:", anchors)
	// Check mode for graphs
	graphmode := "Unlimited Samples"
	if *reset {
		graphmode = "Last 100 Samples"
	}

	t, err := template.ParseFiles("indexTemplate.html")
	if err != nil {
		panic(err)
	}
	timeStr := "Tag " + tagID + " Debug Plots  -  " + (time.Now().Format(time.RFC850))
	titleStr := tagID + " plots "
	descrStr := "Distanses, RSSI and TQF from " + strconv.Itoa(len(anchors)) + " Anchors (100 samples/5s). Mode: " + graphmode
	content := indexContent{Title: titleStr, Header: timeStr, Description: descrStr, Anchors: anchors, Samples: samples}

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

func createTagsIndexPage(TagsList []string) {

	t, err := template.ParseFiles("indexTempTags.html")
	if err != nil {
		panic(err)
	}
	timeStr := " Debug Plots  -  " + (time.Now().Format(time.RFC850))
	titleStr := " Tag plots "
	descrStr := "Select which Tag to display data"
	content := indexContent{Title: titleStr, Header: timeStr, Description: descrStr, Tags: TagsList}

	//create file
	fo, err := os.Create("plots/Tags_Index.html")
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
	err = t.ExecuteTemplate(fo, "indexTempTags.html", content)
	if err != nil {
		panic(err)
	}
}

// Median gets the median number in a slice of numbers
func Median(input []float64) float64 {

	// Start by sorting a copy of the slice
	sort.Float64s(input)

	mid := len(input) / 2
	median := input[mid]
	if len(input)%2 == 0 {
		median = (median + input[mid/2-1]) / 2
	}

	return median
}

type indexContent struct {
	Title       string
	Header      string
	Description string
	Anchors     []string
	Tags        []string
	Samples     map[string]string
}
