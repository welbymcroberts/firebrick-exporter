package collectors

import (
	"encoding/xml"
	"firebrick-exporter/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Threads
/*
-------:/+oo+/::::://+++/++++++//:::.```   `````-/shdmmddddddmddddddmmdyo:.``````````://///os+/////++++////
-----:::/+oo++/+ooosyhddhhhdhhhhhy+.```  ````.:shhhshsyyhhhhdmmysyysyyyyddy+.`````````-++//+so//++++++++///
--::/++osyyhhhhhhhyyhhddmmNNNNNNNh.``` `````:shdddhyhddmmmmmmmmmmmmmdhhysyddy-`````````/+//+so/+ooo+/+++//+
ssyyhhhhhhdddhyyyhhyhhhhdmmNNNNNN/````````.+yhddddddhhyysssssssssyyhhdmmmdmddy-````````/soooss+oooo+++++/++
ddddddhhhddmddhhhddddddhdmmmmmmNd-```````.oyhhysoo+////:::::::::::////+osyhddds-```````smddsoydddhhhyyyysos
ddmmdhhhdmmmmdddddddddddmmmmmmmmd.```  `.oyy+/::----:::://////////:::::::::/ohho``````.hmmmysydmmmmdddmmddd
dddddhhhddddhhhyyyssssssyhhhhhhhh-```  `+o/..-..-::::-------------://///:....-/s:`````/hdddssyhdmmddhdhhhdd
hhhyyyssssoo++/::--:::::///++++++:``` `./-`..-...```````...........```..-----.`..````.syyyysoshddddddhyyhmm
o++//:///:-........-:::::////////+:```-.--.````....-::::::::--:::::--........--...```:ooooooosssyyyhyyysyyd
------/++/-...-----:::--://///////+/-.o/.`.--:::++/ooooossssooso++++//////:-----.s-`-///////oo+++ooo++++///
:////+oso+::::::::::::--:////////////:..`:ysooosyyyyyyssyyhhddmddhso+++oooosooss.--:+ooso/:///:/osssoo++///
++++++ooo+:-........-:--://////////////-.yhyhhddhyyyyyyyyyhdmNNNNNmdysssossyyhhd+///++oo+::::::/ossyso++osy
----.-:oo+:........-:::::///////////////odhddhdhhhhdddmmmmmmmmmmNNNNmmmdhyyhhdddy+++/://:----::/+osyo+//+oy
......:+so:...------:::::///////////////yddhyhhdmmNNNNNNNNmmmddddmNNNNNNmddhhdmms///////:----:://osyo+/::::
-----.-+oo/...---:::::::///::::---::////+shsyyddmddddddmmmddhhdddmdhhddmNNNmdddm+::-///:------://+sys+///++
/:::--:+ss+:-----:--:::////:--...-::://:::syyyhdhs+++//oydmmmmddmd/----:/odNNmmd//+++/:--.-.-::://syso+oooo
::::--/+oo+/:::--...------::::-..-::-----:ohhddho///++//+hyhNmyymh:::-:/:/ymNmmd//osoo+:---.-://:/+o++++//+
+++/::/++////:-.......``...------:///:-..-/yhyysyysyhhhddyomNNhhNNho++ssydddmmmh:/oo+//:---.-:::--:://:...-
+++///++:-----.`.---..``..`````..-:///:-.-:+yyddmNNNNNmds/sNNmmmmNNNhhmNNmmdhdmh-:/++/:------:/:-.:-::-````
::::://:-..`.....-----...`````````..-::----odmmdmmNmddsoydmNNNmmmNmmdmmmNmdddddh:--::-..---.-//::::--:-.```
--::/++/-...-::.......-----...```````.....:ommdysyyssydmNNmmNNNNNmddhhyhhhhhhddh:--..``.-...:+:::-:---:-```
-:/+oso/-..-::-..````...--:::-``````..````-ohdhysssyydmmmmmmmmmddhhyyhhyyyysyhs+....```.....---::`::::::-.`
.-/+ooo/..-::-....--..````.....````.....``.+ydddhddddddhyo///:://////+syhyysys/-.``````.-.---:-::.-+o+/::-`
`-/+++/:--::::--:::-.`````````````........oddmdhhdddyso/:::::::///+++++oyyyyo-.````````.-----:-::::oys+:-.`
.:///+//:::-------.````..``````````````.:/ossyo+ooso+++/::::////:///////+oo+-.`````````.-------::::+sss+-.`
.-:://::::-..````````...``````````````./+/:-:-.:////+//+ooo+++//++ooo/:/+::-----..`...`.-------::::/ooo+/-`
.-:///:-...`````````...`````````````./++/::-.``///::::+++//://://///+//::-`..--:-..````.-----:-::://oooo+:`
.--://-.````````````.``````````````:/+/:::::.``-+/:--://////+++++/////:-.` ``..-...```..-----/-:::::+ooo+:`
--::::.```````....```````````   `-:::::--::--.``/+/-.-::/++osso+ooo+/::/`  ``.....```.o:----::-:::/:ooo+/.`
.----..``````.....````````.-/-.`./:-.----:::--.``:o+:...-::://:::::--/s/`  ```....``  `.---:::--::::+o++/.`
```.````   ``.---.```..-::smNo-`:::-:----::::/:.``:+o+:-..`..```..-:+yy-`  ```.....``  `---:::.-::::sooys:`
  ```      `.-::-+/-:/++//++/.`.:/:-:::::-----...``-+oo+/--.....---/yys.`  ``......``  `---:::--::::sdddmd-
`````   ````.-:+ymNd+::-::--...-//:::--::----....```./ooss+/------+hhy+`` ```...`..``` `----::--:::::-:/++.
```````````-+syddmNNdo-...-//-.///:---::---:-....`````-osyyyyo/--ohhhy:` `````......`` `-----:..---:-:::-.`
`````````-syydmmNmdyo+/+o/++/--o+/+///:--:/:..`.`.``..`./syys:````-oys.````````.....`` `--:--:..://-.-:::::
````````sdysdmmhsoooo+/+//++:-:o++++/:---:-..`.```.....``-o/````````-/``````````````````--:--:.`-//-----:::
.`````.+ssdmyo/+ossso:++/:::--:o+//:::----...`````.`.....``` ```````````````````` ``.```------.`.:::/:--..`
``````./+:./:+osssss/+s+:----.:o/:----:---...`` ```````....``  ``````` `.``````.```  ```---:--. .::.-.`````
*/
type ThreadStat struct {
	XMLName        xml.Name `xml:"stats"`
	Name           string   `xml:"name,attr"`
	Ticks          uint64   `xml:"ticks,attr"`
	CPUPercent     uint8    `xml:"cpu-percent,attr,omitempty"`
	State          string   `xml:"state,attr,omitempty"`
	Lock           string   `xml:"lock,attr,omitempty"`
	Stack          uint64   `xml:"stack,attr,omitempty"`
	Used           uint64   `xml:"used,attr,omitempty"`
	Current        uint64   `xml:"curr,attr,omitempty"`
	CurrentPercent uint8    `xml:"curr-percent,attr,omitempty"`
	DynamicMem     uint64   `xml:"dynamic-mem,attr,omitempty"`
	Priority       int32    `xml:"dynamicpriority,attr,omitempty"`
	Slice          uint64   `xml:"slice,attr,omitempty"`
}
type Threads struct {
	XMLName xml.Name     `xml:"status"`
	Threads []ThreadStat `xml:"thread>stats"`
}

func CollectThreads(device *config.DeviceConfig) {
	// create Client with 3 Second Timeout
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	response, err := httpClient.Get(fmt.Sprintf("http://%s:%s@%s%s", device.Username, device.Password, device.Address, "/status/system/threads?xml"))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	t := Threads{}
	if err := xml.Unmarshal(body, &t); err != nil {
		log.Fatalln(err)
	}
	for _, s := range t.Threads {
		_ = s
		//fmt.Printf("Thread %s: %s%% CPU\n", s.Name, s.Ticks)
	}
}
