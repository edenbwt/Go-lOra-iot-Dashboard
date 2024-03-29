package mqtt

import (
	"database/sql"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	latestReading Message
	mu            sync.RWMutex // Mutex pour accéder en toute sécurité à latestReading de manière concurrente
)

// InitializeDB initialise la connexion à la base de données MariaDB
func InitializeDB() (*sql.DB, error) {
	// Remplacez les paramètres de connexion par vos identifiants MariaDB
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/lora")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// LogToDatabase insère le message reçu dans la base de données MariaDB
func LogToDatabase(db *sql.DB, message Message) error {
	// Préparer l'instruction SQL
	stmt, err := db.Prepare("INSERT INTO data (time, temps_min, temps_curent, temps_max) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Exécuter l'instruction SQL avec les données du message
	_, err = stmt.Exec(message.ReceivedAt, message.UplinkMessage.TemperatureData.Temperature1, message.UplinkMessage.TemperatureData.Temperature2, message.UplinkMessage.TemperatureData.Temperature2)
	if err != nil {
		return err
	}
	return nil
}

// SetLatestReading définit la dernière lecture MQTT
func SetLatestReading(reading Message) {
	mu.Lock()
	defer mu.Unlock()
	latestReading = reading
}

// GetLatestReading renvoie la dernière lecture MQTT
func GetLatestReading() Message {
	mu.RLock()
	defer mu.RUnlock()
	return latestReading
}

func StartMQTT() {
	// Configurer les options du client MQTT
	opts := MQTT.NewClientOptions().AddBroker("tcp://eu1.cloud.thethings.network:1883")
	opts.SetClientID("70B3D54994C49724")
	opts.SetUsername("temps-ciel@ttn")
	opts.SetPassword("NNSXS.4RVYYWM7JGBVBS5XLJZR2OZBLRZWFIMO7KOURSY.7ECNMVZIHJCSWUO6YA5ANMF6WPGOQZEJC6GE6BCOCQ3SP27CM3TA")

	// Créer un nouveau client MQTT
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	// Configurer le canal pour recevoir les signaux de fermeture en douceur
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// S'abonner à un sujet
	topic := "#" // Remplacez your_topic_here par le vrai sujet
	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Printf("Abonné au sujet : %s\n", topic)

	// Attendre les signaux
	<-sigChan
	fmt.Println("Arrêt en cours...")
}

// Définir la fonction de gestionnaire de messages
func messageHandler(client MQTT.Client, msg MQTT.Message) {
	// Analyser la charge utile du message reçu dans la structure
	var receivedMessage Message
	err := json.Unmarshal(msg.Payload(), &receivedMessage)
	if err != nil {
		log.Println("Erreur d'analyse du message :", err)
		return
	}

	// Définir la dernière lecture
	SetLatestReading(receivedMessage)
	log.Println("[§] données récupérées depuis MQTT")

	// Journaliser le message reçu dans la base de données
	db, err := InitializeDB()
	if err != nil {
		log.Println("Erreur de connexion à la base de données :", err)
		return
	}
	defer db.Close()

	err = LogToDatabase(db, receivedMessage)
	if err != nil {
		log.Println("Erreur de journalisation dans la base de données :", err)
		return
	}
	log.Println("Données correctement journalisées dans la base de données")
}
