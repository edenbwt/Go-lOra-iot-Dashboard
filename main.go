package main

import (
	"main.go/lib/mqtt"
	"main.go/lib/webser"
)

func main() {
	go mqtt.GenerateRandomTemperatures()
	webser.StartServ()

}
