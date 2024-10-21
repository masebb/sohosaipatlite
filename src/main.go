package main

import (
	"net"
	"sohosaiPatlite/src/sohosaiPatliteLib"
	"time"
)

func main() {
	p := sohosaiPatliteLib.Patlite{
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

var Off = sohosaiPatliteLib.PatlitePattern{Name: "off", ID: 0}
var On = sohosaiPatliteLib.PatlitePattern{Name: "on", ID: 1}
