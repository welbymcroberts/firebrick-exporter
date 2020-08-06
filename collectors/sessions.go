package collectors

import (
	"encoding/xml"
	"firebrick-exporter/config"
	"log"
)

// Sessions
type Session struct {
	Protocol uint8 `xml:"protocol,attr,omitempty"`
}
type Sessions struct {
	XMLName xml.Name  `xml:"status"`
	Forward []Session `xml:"sessions>forward"`
	Reverse []Session `xml:"sessions>reverse"`
}

func CollectSessions(device *config.DeviceConfig) {
	body, err := CollectViaHTTP(device, "/status/sessions/xml")
	if err != nil {
		log.Fatalln(err)
	}
	s := Sessions{}
	if err := xml.Unmarshal(body, &s); err != nil {
		log.Fatalln(err)
	}
	SessionsTotal := 0
	SessionsTCP := 0
	SessionsUDP := 0
	SessionsICMPv6 := 0
	SessionsICMPv4 := 0
	SessionsVRRP := 0

	// We only check forward sessions, as there will (usually) be a session for both ways
	for _, sess := range s.Forward {
		SessionsTotal++

		if sess.Protocol == 6 {
			SessionsTCP++
		}
		if sess.Protocol == 17 {
			SessionsUDP++
		}
		if sess.Protocol == 112 {
			SessionsVRRP++
		}
		if sess.Protocol == 1 {
			SessionsICMPv4++
		}
		if sess.Protocol == 58 {
			SessionsICMPv6++
		}
	}
	/* fmt.Printf("Sessions: %d, TCP: %d, UDP: %d, VRRP: %d, ICMPv4: %d, ICMPv6: %d\n", SessionsTotal,
	SessionsTCP,
	SessionsUDP,
	SessionsVRRP,
	SessionsICMPv4,
	SessionsICMPv6)
	*/
}
