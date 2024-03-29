#####################################################
## Liste des librairies utilisées
######################################################

from machine import Timer
from machine import I2C
# configure the I2C bus
import pycom
from network import LoRa
import socket
import time
import ubinascii
from machine import UART
import machine
import math

#####################################################
## PENSER A LIRE LES COMMENTAIRES
######################################################

##Init variables
periodeHorloge = 180
tempMax = -100
tempMin = 100
address = 24
temp_reg = 5
res_reg = 8
dataAvailable = 0
pycom.heartbeat(False)


# Creation de la trame au format Cayenne
# TRAME TEMPORAIRE. IL FAUT EN CREER UNE AU FORMAT CAYENNE POUR 

maTrame = [
    0x01, 0x67, 0x00, 0x00,  # Température actuelle
    0x02, 0x67, 0x00, 0x00,  # Température maximale
    0x03, 0x67, 0x00, 0x00   # Température minimale
]

######################################################
## Init du capteur de temperature
######################################################

i2c = I2C(0, I2C.MASTER, baudrate=100000)
i2c.scan() # returns list of slave addresses

def temp_c(data):
    value = data[0] << 8 | data[1]
    temp = (value & 0xFFF) / 16.0
    if value & 0x1000:
        temp -= 256.0
    return temp

##################################Fin

######################################################
## Creation d'un objet timer
######################################################

class Clock:
    #Constructeur
    def __init__(self):
        #Init timer
        self.num_paquet = 0
        #self.__alarm = Timer.Alarm(Fct a executer toute, temps entre chaque execution , periodic ou Non)
        self.__alarm = Timer.Alarm(self._seconds_handler, periodeHorloge, periodic=True)

    #Fonction a exectuer a chaque tic du timer
    def _seconds_handler(self, alarm):
        # Incrémentation du numéro de paquet à chaque appel
        self.num_paquet += 1
        global tempMax
        global tempMin
        
        # Lecture de la température actuelle depuis le capteur via I2C
        data = i2c.readfrom_mem(address, temp_reg, 2)
        tempCurrent = temp_c(data)
        
        # Conversion et mise à jour de la température actuelle dans la trame maTrame
        temp_current_value = int(tempCurrent * 10)
        maTrame[2] = (temp_current_value >> 8) & 0xFF  # High byte
        maTrame[3] = temp_current_value & 0xFF         # Low byte

        # Conversion et mise à jour de la température maximale dans la trame maTrame
        temp_max_value = int(tempMax * 10)
        maTrame[6] = (temp_max_value >> 8) & 0xFF  # High byte
        maTrame[7] = temp_max_value & 0xFF         # Low byte

        # Conversion et mise à jour de la température minimale dans la trame maTrame
        temp_min_value = int(tempMin * 10)
        maTrame[10] = (temp_min_value >> 8) & 0xFF  # High byte
        maTrame[11] = temp_min_value & 0xFF         # Low byte

        # Envoi de la trame maTrame via LoRa
        s.send(bytes(maTrame))

        #Possibilite d'arreter le timer
        #if self.seconds == 100:
        #    alarm.cancel() # stop counting after 10 seconds

######################################################
## Fin timer
######################################################



######################################################
## Debut Programme
######################################################

#Creation du Timer
clock = Clock()

######################################### Init du LoRa.
# Selectionner la bonne région
# Asia = LoRa.AS923
# Australia = LoRa.AU915
# Europe = LoRa.EU868
# United States = LoRa.US915
lora = LoRa(mode=LoRa.LORAWAN, region=LoRa.EU868)

# Creez des parametres d'authentification OTAA. Remplacez-les par les informations presentes sur TTN
# A MODIFIER EN FONCTION DE VOS DONNEES TTN
app_eui = ubinascii.unhexlify('70B3D54994C49724')
app_key = ubinascii.unhexlify('9DEEC59D0CD308145EB7F8B71D2A3D76')
dev_eui = ubinascii.unhexlify('70B3D57ED0066239')

######################################################
## Fonction lorsqu'un evennement lora est detecte
######################################################

def lora_cb(lora):
    global s
    events = lora.events()
    #Si Reception de donnée 'Downlink'
    if events & LoRa.RX_PACKET_EVENT:
        print('Packet recu')
        message = list(s.recv(11))
        #En fonction de la valeur du dernier octets
        #La couleur des leds change
        maTrame[10] = message[10]
        if maTrame[10] == 1 :
            pycom.rgbled(0xFF0000)  # Red
        elif maTrame[10] == 2 : 
            pycom.rgbled(0x00FF00)  # Green
        elif maTrame[10] == 3 : 
            pycom.rgbled(0x0000FF)  # Blue
        else :
            pycom.rgbled(0x000000)  
    #ACK de transmission de donnée
    if events & LoRa.TX_PACKET_EVENT:
        print('packet envoye')
        pycom.rgbled(0xFF0000)  # Red
        for e in range(1,10) :
            time.sleep(0.2)
            pycom.rgbled(0xFF0000)  # Green
            time.sleep(0.2)
            pycom.rgbled(0x000000)  # Blue
        
        pycom.rgbled(0x000000)  # Blue

##################################Fin

#Definition de la fonction a executer lorsqu'il y a une tranmssion ou une reception 
lora.callback(trigger=(LoRa.RX_PACKET_EVENT | LoRa.TX_PACKET_EVENT), handler=lora_cb)


#Démarrate de la connexion avec le server Lora
lora.join(activation=LoRa.OTAA, auth=(dev_eui, app_eui, app_key), timeout=0)

#Attente de la connexion
while not lora.has_joined():
    time.sleep(2.5)
    print('Pas encore de connexion')
pycom.rgbled(0xFF0000)  # Red
time.sleep(1)
pycom.rgbled(0x00FF00)  # Green
time.sleep(1)
pycom.rgbled(0x0000FF)  # Blue
time.sleep(1)
pycom.rgbled(0x000000)  # Blue
print('Youpi mon device a rejoint un serveur')


#Creation du'une socket Lora
s = socket.socket(socket.AF_LORA, socket.SOCK_RAW)

# Regelage de la connexion 0 POUR SF = 12 Meilleur distance
s.setsockopt(socket.SOL_LORA, socket.SO_DR, 0)

# Rendre la socket bloquante
#NE PAS Attente que les données soient envoyées 
s.setblocking(False)

#Boucle infinie 
#Elle s'arrete uniquement lorsqu'il y a une interrutpion du timer
#et reprend après
while 1 : 
    data = i2c.readfrom_mem(address, temp_reg, 2)
    tempCurrent = temp_c(data)
    
    if tempCurrent > tempMax : 
        tempMax = tempCurrent
    if tempCurrent < tempMin : 
        tempMin = tempCurrent
   
#####################################################
## FIN
######################################################


