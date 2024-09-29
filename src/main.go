package main

import (
	"sohosaiPatlite/src/sohosaiPatliteLib"
	"time"
)

func main() {
	sohosaiPatliteLib.NewPatlite("192.168.10.1")
	for {
		sohosaiPatliteLib.SetPatlite(false, false, false, false)
		time.Sleep(500 * time.Millisecond)
		sohosaiPatliteLib.SetPatlite(true, false, false, false)
		time.Sleep(100 * time.Millisecond)
		sohosaiPatliteLib.SetPatlite(false, true, false, false)
		time.Sleep(100 * time.Millisecond)
		sohosaiPatliteLib.SetPatlite(false, false, true, false)
		time.Sleep(100 * time.Millisecond)
		sohosaiPatliteLib.SetPatlite(false, false, false, false)
		time.Sleep(500 * time.Millisecond)
		sohosaiPatliteLib.SetPatlite(false, false, true, false)
		time.Sleep(100 * time.Millisecond)
		sohosaiPatliteLib.SetPatlite(false, true, false, false)
		time.Sleep(100 * time.Millisecond)
		sohosaiPatliteLib.SetPatlite(true, false, false, false)
		time.Sleep(100 * time.Millisecond)
	}
}
