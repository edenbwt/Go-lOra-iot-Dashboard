###################################################################
### README - Programme de Surveillance de Température avec LoRa ###
###################################################################

########## Description ##########

Ce programme est conçu pour surveiller la température à l'aide d'un capteur de température connecté via I2C et envoyer les données de température actuelle, maximale et minimale à un serveur LoRaWAN en utilisant le protocole LoRa. Les données sont envoyées sous forme de trame Cayenne.

########## Prérequis ##########

Carte de développement compatible avec MicroPython (par exemple, FiPy)
Capteur de température compatible avec le programme (capteur utilisant l'interface I2C)
Connexion à un réseau LoRaWAN


########## Installation et Configuration ###########

##### Bibliothèques nécessaires #####
Le programme utilise plusieurs bibliothèques MicroPython. Assurez-vous de les installer sur votre carte de développement avant d'exécuter le programme :

machine
pycom
network
socket
time
ubinascii
math

##### Configuration du LoRa #####

Avant d'exécuter le programme, vous devez configurer les paramètres LoRa pour votre région et les informations d'authentification OTAA :

Sélectionnez la région LoRa appropriée (Europe, United States, Asia, etc.)
Remplacez app_eui, app_key, et dev_eui par les informations d'authentification fournies par votre fournisseur LoRaWAN.

##### Configuration du Capteur de Température #####
Le programme est configuré pour utiliser un capteur de température connecté via I2C. Assurez-vous que l'adresse du capteur (address), le registre de température (temp_reg), et le registre de résolution (res_reg) sont correctement définis pour votre capteur.

########## Utilisation ##########

Téléchargez et installez les bibliothèques nécessaires sur votre carte de développement.
Configurez les paramètres LoRa et les informations d'authentification OTAA dans le programme.
Connectez votre capteur de température à la carte de développement et assurez-vous qu'il est correctement configuré.
Téléchargez et exécutez le programme sur votre carte de développement.

########## Fonctionnalités du Programme ##########

##### Surveillance de la Température #####
Capteur de Température : Utilise le capteur de température connecté via I2C pour lire la température actuelle à intervalles réguliers.
Mise à Jour des Valeurs : Met à jour les valeurs de température maximale et minimale en fonction de la température actuelle lue par le capteur.

##### Envoi des Données via LoRa #####
Trame Cayenne : Les données de température sont formatées et envoyées sous forme de trame Cayenne pour une transmission efficace.
Transmission LoRa : Utilise le protocole LoRa pour envoyer les données de température au serveur LoRaWAN configuré.

##### Gestion des Événements LoRa #####
Événements LoRa : Gère les événements LoRa tels que la réception de données 'Downlink' ou l'ACK de transmission de données.
LEDs Indicatrices : Change la couleur des LEDs (Rouge, Vert, Bleu) en fonction de la valeur du dernier octet reçu pour indiquer le statut de la réception des données.

