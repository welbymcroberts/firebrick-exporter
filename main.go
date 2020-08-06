package main

import (
	"firebrick-exporter/collectors"
	"firebrick-exporter/config"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"reflect"
	"sync"
	"time"
)

var conf *config.Config

func init() {
	flag.Usage = func() {
		flag.PrintDefaults()
	}
}

func main() {
	// Parse command line flags
	flag.Parse()

	// Load Config
	c, err := config.ConfigLoadFromFile()
	if err != nil {
		log.Fatal(err)
	}
	conf = c

	// We're going against some of the recommended practices for prometheus exporters
	// where the scrape of metrics is done at a different time to the scrap from prometheus
	// this is due to there being a possibility of the FB's being polled having a high latency
	// low reliability link, which would potentially mean that there is an increase in polling time
	// To combat this, we're going to put a last scrape time, and a success/failure status
	// for each of the devices
	ticker := time.NewTicker(10 * time.Second)
	done := make(chan bool)

	// do the actual collection - Currently this may overlap with a previous run
	// TODO: Check that we dont have another tick running
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				start := time.Now()
				collectByDevice()
				elapsed := time.Since(start)
				fmt.Printf("Scrape of all devices took %s\n", elapsed)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	// TODO: Set port to something that's registered on Prom Wiki
	http.ListenAndServe(":60666", nil)
}

func collectByDevice() {
	deviceWG := sync.WaitGroup{}
	deviceWG.Add(len(conf.Devices))

	for _, device := range conf.Devices {
		go func() {
			// TODO: check that the http endpoint is reachable is running, fail if not
			// Create waitgroup for all enabled tests
			wg := sync.WaitGroup{}

			// Get number of enabled features
			val := reflect.ValueOf(device.Features)
			for i := 0; i < val.NumField(); i++ {
				if val.Field(i).Interface() == true {
					wg.Add(1)
				}
			}

			// Check Each feature
			// TODO: Consider how we handle failures. per Collector, or per device?
			if device.Features.Power == true {
				collectors.CollectPower(device)
				wg.Done()
			}
			if device.Features.Threads == true {
				collectors.CollectThreads(device)
				wg.Done()
			}
			if device.Features.Subnets == true {
				collectors.CollectSubnets(device)
				wg.Done()
			}
			if device.Features.DNS == true {
				collectors.CollectDNS(device)
				wg.Done()
			}
			if device.Features.DHCP == true {
				collectors.CollectDHCP(device)
				wg.Done()
			}
			if device.Features.Profiles == true {
				collectors.CollectProfiles(device)
				wg.Done()
			}
			if device.Features.PPPoE == true {
				collectors.CollectPPPoE(device)
				wg.Done()
			}
			if device.Features.L2TP == true {
				//collectL2TP(device)
				wg.Done()
			}
			if device.Features.Sessions == true {
				collectors.CollectSessions(device)
				wg.Done()
			}
			// Wait until everything is completed
			wg.Wait()
			fmt.Printf("%s - done\n", device.Address)
			deviceWG.Done()
		}()
	}
	deviceWG.Wait()
}
