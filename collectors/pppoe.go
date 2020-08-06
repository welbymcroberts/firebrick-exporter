package collectors

import (
	"encoding/xml"
	"firebrick-exporter/config"
	"log"
)

// PPPoE
type PPPoE struct {
	Name         string `xml:"name,attr"`
	Status       string `xml:"status,attr"`
	Port         string `xml:"port,attr,omitempty"`
	RoutingTable uint16 `xml:"table,attr,omitempty"`
	ID           uint16 `xml:"ID,attr,omitempty"`
	LocalID      string `xml:"local-id,attr,omitempty"`
	RemoteID     string `xml:"remote-id,attr,omitempty"`
	Response     string `xml:"response,attr,omitempty"`
	MTU          uint16 `xml:"mtu,attr,omitempty"`
}

type PPPoEs struct {
	XMLName  xml.Name `xml:"status"`
	Active   []PPPoE  `xml:"pppoe>active"`
	Inactive []PPPoE  `xml:"pppoe>try"`
}

func CollectPPPoE(device *config.DeviceConfig) {

	body, err := CollectViaHTTP(device, "/status/pppoe/xml")
	if err != nil {
		log.Fatalln(err)
	}

	pppoe := PPPoEs{}
	if err := xml.Unmarshal(body, &pppoe); err != nil {
		log.Fatalln(err)
	}
	for _, in := range pppoe.Inactive {
		//fmt.Println(in)
		_ = in
	}
	for _, a := range pppoe.Active {
		// fmt.Println(a)
		_ = a
	}

}
