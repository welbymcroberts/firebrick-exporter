package collectors

import (
	"encoding/xml"
	"firebrick-exporter/config"
	"log"
)

// Profiles
type Profile struct {
	Name   string `xml:"name,attr"`
	Status string `xml:"status,attr"`
	Reason string `xml:"reason,attr,omitempty"`
}

type Profiles struct {
	XMLName  xml.Name  `xml:"status"`
	Active   []Profile `xml:"profiles>active"`
	Inactive []Profile `xml:"profiles>inactive"`
}

func CollectProfiles(device *config.DeviceConfig) {

	body, err := CollectViaHTTP(device, "/status/profiles/xml")
	if err != nil {
		log.Fatalln(err)
	}

	profiles := Profiles{}
	if err := xml.Unmarshal(body, &profiles); err != nil {
		log.Fatalln(err)
	}
	for _, in := range profiles.Inactive {
		//fmt.Println(in)
		_ = in
	}
	for _, a := range profiles.Active {
		//fmt.Println(a)
		_ = a
	}

}
