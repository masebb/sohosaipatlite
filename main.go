package sopolight

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"time"
)

func main() {
	p := Patlite{
		Ipaddr: net.ParseIP("192.168.10.1"),
		Red:    On,
		Yellow: On,
		Green:  On,
	}
	for {
		p.SetPatlite(On, Off, Off, Off)
		time.Sleep(500 * time.Millisecond)
		p.SetPatlite(Off, On, Off, Off)
		time.Sleep(100 * time.Millisecond)
		p.SetPatlite(Off, Off, On, Off)
		time.Sleep(100 * time.Millisecond)
		p.SetPatlite(Off, Off, Off, On)
		time.Sleep(100 * time.Millisecond)
		p.SetPatlite(Off, Off, Off, Off)
		time.Sleep(500 * time.Millisecond)
		p.SetPatlite(On, On, On, On)
		time.Sleep(100 * time.Millisecond)
		p.SetPatlite(Off, Off, Off, Off)
		time.Sleep(100 * time.Millisecond)
		p.SetPatlite(On, On, On, On)
		time.Sleep(100 * time.Millisecond)
	}
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

func (p *Patlite) SetPatlite(red PatlitePattern, yellow PatlitePattern, green PatlitePattern, buzzer PatlitePattern) {
	p.Red = red
	p.Yellow = yellow
	p.Green = green
	p.Buzzer = buzzer
	p.sendAlertToPatlite()
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

type PatlitePattern struct {
	Name string
	ID   int
}

func (p *Patlite) sendAlertToPatlite() {
	url := "http://" +
		p.Ipaddr.String() +
		"/api/control?alert=" +
		strconv.Itoa(p.Red.ID) +
		strconv.Itoa(p.Yellow.ID) +
		strconv.Itoa(p.Green.ID) +
		"0" + // 青は未使用
		"0" + // 白は未使用
		strconv.Itoa(p.Buzzer.ID)

	log.Println("GET request to ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error at set: ", err)
	}
	defer resp.Body.Close()

	log.Println(resp)
}
