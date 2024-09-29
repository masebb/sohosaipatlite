package sohosaiPatliteLib

import (
	"log"
	"net"
	"net/http"
)

type patlite struct {
	Ipaddr net.IP
	Red    bool // 実はモードがあるのでintにしたほうがいいかも
	Yellow bool
	Green  bool
	Buzzer bool
}

var p patlite

func NewPatlite(ip string) {
	p = patlite{
		Ipaddr: net.ParseIP(ip),
		Red:    false,
		Yellow: false,
		Green:  false,
		Buzzer: false,
	}
	sendAlertToPatlite()
}

func SetPatlite(red bool, yellow bool, green bool, buzzer bool) {
	p.Red = red
	p.Yellow = yellow
	p.Green = green
	p.Buzzer = buzzer

	sendAlertToPatlite()
}
func sendAlertToPatlite() {
	url := "http://" +
		p.Ipaddr.String() +
		"/api/control?alert=" +
		boolToIntString(p.Red) +
		boolToIntString(p.Yellow) +
		boolToIntString(p.Green) +
		"0" + // 青は未使用
		"0" + // 白は未使用
		boolToIntString(p.Buzzer)

	log.Println("GET request to ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error at set: ", err)
	}
	defer resp.Body.Close()

	log.Println(resp)
}

func SendClearToPatlite() {
	url := "http://" +
		p.Ipaddr.String() +
		"/api/control?clear=1"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error at clear: ", err)
	}
	defer resp.Body.Close()
}

func boolToIntString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
