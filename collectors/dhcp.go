package collectors

import (
	"encoding/xml"
	"firebrick-exporter/config"
	"log"
)

// DHCP
type DHCPLease struct {
	Port    string `xml:"port,attr"`
	Expired bool   `xml:"expired,attr,omitempty"`
}

type DHCPLeases struct {
	XMLName  xml.Name    `xml:"status"`
	Active   []DHCPLease `xml:"dhcp>active"`
	Inactive []DHCPLease `xml:"dhcp>inactive"`
}

func CollectDHCP(device *config.DeviceConfig) {

	body, err := CollectViaHTTP(device, "/status/dhcp/xml")
	if err != nil {
		log.Fatalln(err)
	}

	dhcp := DHCPLeases{}
	if err := xml.Unmarshal(body, &dhcp); err != nil {
		log.Fatalln(err)
	}
	for _, in := range dhcp.Inactive {
		//fmt.Println(in)
		_ = in
	}
	for _, a := range dhcp.Active {
		//fmt.Println(a)
		_ = a
	}

}
