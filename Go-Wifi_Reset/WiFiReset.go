package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tatsushid/go-fastping"
)

func main() {
	fmt.Println("It hath Begun")
	for {
		addr := "8.8.8.8"
		ping := fastping.NewPinger()
		ping.Source("192.168.0.176")  //192.168.20.240
		ping.AddIP(addr)
		ping.RunLoop()
		ping.OnRecv(){
		time.Sleep(1000*time.Millisecond)
	}

}

func reset() {
	cmnd := exec.Command("WiFi_Reset.bat")
	cmnd.Start()
	t := time.Now().String()
	log.Printf("Wifi reset at %s ", t)
}
