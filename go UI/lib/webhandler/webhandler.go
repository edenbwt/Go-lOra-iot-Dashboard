package webhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main.go/lib/mqtt"
	"net/http"
)

// Structure de réponse pour les lectures de température
type TemperatureReadingResponse struct {
	Time          string                            `json:"time"`           // Temps de réception de la lecture
	EndDeviceIDs  mqtt.EndDeviceIDs                 `json:"end_device_ids"` // Identifiants de l'appareil final
	UplinkMessage mqtt.UplinkMessageWithTemperature `json:"uplink_message"` // Message montant avec la température
	NetworkIDs    mqtt.NetworkIDs                   `json:"network_ids"`    // Identifiants du réseau
}

// Gestionnaire de requêtes pour récupérer les lectures de température
func GetTemperatureReadingHandler(w http.ResponseWriter, r *http.Request) {
	// Obtenir la dernière lecture de température depuis MQTT
	latestReading := mqtt.GetLatestReading()

	// Convertir le temps au format désiré
	receivedAt := latestReading.ReceivedAt.Format("2006-01-02 15:04:05")

	// Créer une structure de réponse
	response := TemperatureReadingResponse{
		Time:          receivedAt,
		EndDeviceIDs:  latestReading.EndDeviceIDs,
		UplinkMessage: latestReading.UplinkMessage,
		NetworkIDs:    latestReading.NetworkIDs,
	}

	// Marshaler la réponse en JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Échec du marshalling JSON", http.StatusInternalServerError)
		return
	}

	// Définir l'en-tête de réponse
	w.Header().Set("Content-Type", "application/json")

	// Écrire la réponse JSON
	w.Write(jsonData)
}

// Gestionnaire pour le capteur Web
func HandleWebSensor(w http.ResponseWriter, r *http.Request) {
	// Lire le fichier HTML
	htmlContent, err := ioutil.ReadFile("C:\\Users\\SNIR_admin\\Desktop\\Go Lora\\LOra Project\\lib\\webhandler\\sensor.html")
	if err != nil {
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Définir l'en-tête de réponse
	w.Header().Set("Content-Type", "text/html")

	// Écrire la réponse HTML
	fmt.Fprintf(w, "%s", htmlContent)
}
