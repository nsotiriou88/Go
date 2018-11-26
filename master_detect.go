package MasterDetect

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/grandcat/zeroconf"
)

// MasterDetect detects the Master throw mDNS(Bonjour) and returns its
// ID(uint16) and IP(net.IP).
func MasterDetect() {
	var MasterID uint16
	var MasterIP net.IP

	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			if strings.Contains(entry.ServiceRecord.Instance, "Primary") {
				log.Println("Master:", entry.ServiceRecord.Instance[9:13], entry.AddrIPv4[0])
				tempID, err := strconv.ParseUint(entry.ServiceRecord.Instance[9:13], 16, 16)
				if err != nil {
					log.Fatalln("Failed to convert Master's ID:", err.Error())
				}
				MasterID = uint16(tempID)
				MasterIP = entry.AddrIPv4[0]
			}
		}
		log.Println("No more entries.")
	}(entries)

	// One second only is enough to detect Master.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	err = resolver.Browse(ctx, "_svc_openrtls._tcp", "local.", entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()

	return MasterID, MasterIP
}
