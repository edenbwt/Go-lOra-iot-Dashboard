package webser

import (
	"fmt"
	"main.go/lib/webhandler"
	"net/http"
)

// StartServ démarre le serveur Web et enregistre les gestionnaires de route associés.
func StartServ() {
	// Enregistrer la fonction de gestionnaire pour la route "/sensor"
	http.HandleFunc("/sensor", webhandler.HandleWebSensor) // démarrage du gestionnaire pour la page GUI
	// Enregistrer l'API pour récupérer les données de température fournies par MQTT
	http.HandleFunc("/api/temperature", webhandler.GetTemperatureReadingHandler) // démarrage de l'API sur les données collectées par MQTT

	// Démarrer le serveur sur le port 8080, modifiable si nécessaire
	fmt.Println("[!] Serveur démarré sur le port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("[!] Erreur lors du démarrage du serveur :", err)
	}
}
