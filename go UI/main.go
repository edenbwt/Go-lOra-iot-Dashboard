package main

import (
	"main.go/lib/mqtt"
	"main.go/lib/webser"
)

func main() {
	go mqtt.StartMQTT() // demarer separement le server mqtt pour laiser le wenhost ce lancer aussi
	webser.StartServ()  // demarer le server

}
S