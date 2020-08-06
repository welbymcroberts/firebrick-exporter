package collectors

import (
	"encoding/xml"
	"firebrick-exporter/config"
	"log"
)

// DNS

type DNS struct {
	Replied      uint64 `xml:"replied,attr"`
	Sent         uint64 `xml:"sent,attr,omitempty"`
	IP           string `xml:"mtu,attr,omitempty"`
	Queue        uint64 `xml:"queue,attr,omitempty"`
	RoutingTable uint8  `xml:"table,attr,omitempty"`
	// TODO: Timeouts
}

type DNSStats struct {
	XMLName  xml.Name `xml:"status"`
	Active   []DNS    `xml:"dns>active"`
	Inactive []DNS    `xml:"dns>inactive"`
}

func CollectDNS(device *config.DeviceConfig) {

	body, err := CollectViaHTTP(device, "/status/dns/xml")
	if err != nil {
		log.Fatalln(err)
	}

	dns := DNSStats{}
	if err := xml.Unmarshal(body, &dns); err != nil {
		log.Fatalln(err)
	}
	for _, in := range dns.Inactive {
		//fmt.Println(in)
		_ = in
	}
	for _, a := range dns.Active {
		//fmt.Println(a)
		_ = a
	}

}
