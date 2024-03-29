package mqtt

import (
	"time"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  int     `json:"altitude"`
	Source    string  `json:"source"`
}

type GatewayInfo struct {
	GatewayID string `json:"gateway_id"`
	EUI       string `json:"eui"`
}

type RxMetadata struct {
	GatewayIDs  GatewayInfo `json:"gateway_ids"`
	Time        time.Time   `json:"time"`
	Timestamp   int64       `json:"timestamp"`
	RSSI        int         `json:"rssi"`
	ChannelRSSI int         `json:"channel_rssi"`
	SNR         float64     `json:"snr"`
	Location    Location    `json:"location"`
	UplinkToken string      `json:"uplink_token"`
	ReceivedAt  time.Time   `json:"received_at"`
}

type DataRate struct {
	Bandwidth       int    `json:"bandwidth"`
	SpreadingFactor int    `json:"spreading_factor"`
	CodingRate      string `json:"coding_rate"`
}

type Settings struct {
	DataRate  DataRate  `json:"data_rate"`
	Frequency string    `json:"frequency"`
	Timestamp int64     `json:"timestamp"`
	Time      time.Time `json:"time"`
}

type EndDeviceIDs struct {
	DeviceID       string `json:"device_id"`
	ApplicationIDs struct {
		ApplicationID string `json:"application_id"`
	} `json:"application_ids"`
	DevEUI  string `json:"dev_eui"`
	JoinEUI string `json:"join_eui"`
	DevAddr string `json:"dev_addr"`
}

type TemperatureData struct {
	Temperature1 float64 `json:"temperature_1"`
	Temperature2 float64 `json:"temperature_2"`
	Temperature3 float64 `json:"temperature_3"`
}

type UplinkMessageWithTemperature struct {
	SessionKeyID    string          `json:"session_key_id"`
	FPort           int             `json:"f_port"`
	FCnt            int             `json:"f_cnt"`
	FRMPayload      string          `json:"frm_payload"`
	RxMetadata      []RxMetadata    `json:"rx_metadata"`
	Settings        Settings        `json:"settings"`
	ReceivedAt      time.Time       `json:"received_at"`
	ConsumedAirtime string          `json:"consumed_airtime"`
	NetworkIDs      NetworkIDs      `json:"network_ids"`
	TemperatureData TemperatureData `json:"decoded_payload"`
}

type NetworkIDs struct {
	NetID          string `json:"net_id"`
	NSID           string `json:"ns_id"`
	TenantID       string `json:"tenant_id"`
	ClusterID      string `json:"cluster_id"`
	ClusterAddress string `json:"cluster_address"`
}

type Message struct {
	EndDeviceIDs   EndDeviceIDs                 `json:"end_device_ids"`
	CorrelationIDs []string                     `json:"correlation_ids"`
	ReceivedAt     time.Time                    `json:"received_at"`
	UplinkMessage  UplinkMessageWithTemperature `json:"uplink_message"`
	NetworkIDs     NetworkIDs                   `json:"network_ids"`
}
