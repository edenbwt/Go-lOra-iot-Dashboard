package webhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main.go/lib/mqtt"
	"net/http"
)

// HandleWeb function to handle incoming requests and display temperature data
type TemperatureReadingResponse struct {
	Time string `json:"time"`
	Data string `json:"data"`
}

// GetTemperatureReadingHandler function to handle requests to fetch temperature data
func GetTemperatureReadingHandler(w http.ResponseWriter, r *http.Request) {
	// Get the latest temperature reading from MQTT
	latestReading := mqtt.GetLatestReading()

	// Create response struct
	response := TemperatureReadingResponse{
		Time: latestReading.Time.Format("2006-01-02 15:04:05"),
		Data: latestReading.Data,
	}

	// Marshal response to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	w.Write(jsonData)
}
func HandleWebSensor(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file
	htmlContent, err := ioutil.ReadFile("C:\\Users\\smurf\\OneDrive\\Bureau\\LOra Project\\lib\\webhandler\\sensor.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response header
	w.Header().Set("Content-Type", "text/html")

	// Write the HTML response
	fmt.Fprintf(w, "%s", htmlContent)
}

func HandleWebDashBoard(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file
	htmlContent, err := ioutil.ReadFile("C:\\Users\\smurf\\OneDrive\\Bureau\\LOra Project\\lib\\webhandler\\dashboard.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response header
	w.Header().Set("Content-Type", "text/html")

	// Write the HTML response
	fmt.Fprintf(w, "%s", htmlContent)
}
