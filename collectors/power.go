package collectors

// Power
import (
	"encoding/xml"
	"firebrick-exporter/config"
	"log"
)

type Readings struct {
	XMLName xml.Name `xml:"reading"`
	Source  string   `xml:"source,attr"`
	Value   string   `xml:"value,attr"`
	Min     string   `xml:"min,attr,omitempty"`
	Max     string   `xml:"max,attr,omitempty"`
	Spread  string   `xml:"spread,attr,omitempty"`
}
type PowerStatus struct {
	XMLName     xml.Name   `xml:"status"`
	Power       []Readings `xml:"power>reading"`
	Temperature []Readings `xml:"temperature>reading"`
}

func CollectPower(device *config.DeviceConfig) {
	body, err := CollectViaHTTP(device, "/status/system/monitor?xml")
	if err != nil {
		log.Fatalln(err)
	}
	p := PowerStatus{}
	if err := xml.Unmarshal(body, &p); err != nil {
		log.Fatalln(err)
	}
	for _, t := range p.Temperature {
		_ = t
		//fmt.Printf("Temperature - Source %s: %s\n", t.Source, t.Value)
	}
	for _, p := range p.Power {
		_ = p
		//fmt.Printf("Power - Source %s: %s\n", p.Source, p.Value)
	}
}
