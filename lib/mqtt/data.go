package mqtt

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TemperatureReading struct {
	Time time.Time
	Data string
}

var latestReading TemperatureReading
var mutex sync.RWMutex

// Function to generate random temperature readings in the background
func GenerateRandomTemperatures() {
	for {
		mutex.Lock()
		latestReading = TemperatureReading{
			Time: time.Now(),
			Data: generateRandomTemperature(),
		}
		mutex.Unlock()
		time.Sleep(time.Second) // Wait for 1 minute before generating the next reading
	}
}

// Function to generate a random temperature between 0 and 50 degrees Celsius
func generateRandomTemperature() string {
	return fmt.Sprintf("%.2f°C", rand.Float64()*50)
}

// Function to get the latest temperature reading
func GetLatestReading() TemperatureReading {
	mutex.RLock()
	defer mutex.RUnlock()
	return latestReading
}
