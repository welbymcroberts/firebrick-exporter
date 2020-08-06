package collectors

import (
	"encoding/xml"
	"firebrick-exporter/config"
	"log"
)

// Subnets

type Interface struct {
	Name         string `xml:"name,attr"`
	MACAddress   string `xml:"our-mac,attr,omitempty"`
	MTU          uint16 `xml:"mtu,attr,omitempty"`
	Interface    string `xml:"interface,attr"`
	Port         string `xml:"port,attr,omitempty"`
	RoutingTable uint8  `xml:"table,attr,omitempty"`
}
type Active struct {
	Name         string `xml:"name,attr"`
	MACAddress   string `xml:"our-mac,attr,omitempty"`
	MTU          uint16 `xml:"mtu,attr,omitempty"`
	Interface    string `xml:"interface,attr"`
	Port         string `xml:"port,attr,omitempty"`
	RoutingTable uint8  `xml:"table,attr,omitempty"`
	// Expiry string `xml:"expiry,attr,omitempty"`
	DHCPServer  string `xml:"dhcp-server,attr,omitempty"`
	DHCPDns     string `xml:"dhcp-dns,attr,omitempty"`
	DHCPGateway string `xml:"dhcp-gateway,attr,omitempty"`
	ID          string `xml:"ID,attr,omitempty"`
	Comment     string `xml:"comment,atrr,omitempty"`
}
type Subnets struct {
	XMLName    xml.Name    `xml:"status"`
	Interfaces []Interface `xml:"subnets>interface"`
	Active     []Active    `xml:"subnets>active"`
}

func CollectSubnets(device *config.DeviceConfig) {
	body, err := CollectViaHTTP(device, "/status/subnets/xml")
	if err != nil {
		log.Fatalln(err)
	}

	s := Subnets{}
	if err := xml.Unmarshal(body, &s); err != nil {
		log.Fatalln(err)
	}
	for _, in := range s.Interfaces {
		_ = in
		//fmt.Println(in)
	}
	for _, a := range s.Active {
		_ = a
		//fmt.Println(a)
	}

}
