package sopolight

import (
	"log"
	"net"
	"net/http"
	"strconv"
)

type PatlitePattern struct {
	Name string
	ID   int
}

var Off = PatlitePattern{Name: "off", ID: 0}
var On = PatlitePattern{Name: "on", ID: 1}

type Patlite struct {
	Ipaddr net.IP
	Red    PatlitePattern
	Yellow PatlitePattern
	Green  PatlitePattern
	Buzzer PatlitePattern
}

func (p *Patlite) SendClearToPatlite() {
	url := "http://" +
		p.Ipaddr.String() +
		"/api/control?clear=1"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error at clear: ", err)
	}
	defer resp.Body.Close()
}

func (p *Patlite) SendAlertToPatlite() {
	url := "http://" +
		p.Ipaddr.String() +
		"/api/control?alert=" +
		strconv.Itoa(p.Red.ID) +
		strconv.Itoa(p.Yellow.ID) +
		strconv.Itoa(p.Green.ID) +
		"0" + // 青は未使用
		"0" + // 白は未使用
		strconv.Itoa(p.Buzzer.ID)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error at set: ", err)
	}
	defer resp.Body.Close()

	log.Println(resp)
}
