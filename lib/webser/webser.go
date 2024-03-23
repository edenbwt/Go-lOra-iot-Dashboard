package webser

import (
	"fmt"
	"main.go/lib/webhandler"
	"net/http"
)

func StartServ() {
	// Register the handler function for the "/hello" route
	http.HandleFunc("/sensor", webhandler.HandleWebSensor)
	http.HandleFunc("/", webhandler.HandleWebDashBoard)
	http.HandleFunc("/api/temperature", webhandler.GetTemperatureReadingHandler)

	// Start the server on port 8080
	fmt.Println("[!] Server started on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("[!] Error starting server:", err)
	}
}
