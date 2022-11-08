package model

// Config is a type for data config.
type Config struct {
	Network     int         `json:"Network"`
	GSMParams   GSMParams   `json:"GsmParams"`
	WCDMAParams WCDMAParams `json:"WcdmaParams"`
	LTEParams   LTEParams   `json:"Lteparams"`
	Slot        string      `json:"Slot"`
	Version     uint32      `json:"Version"`
}

// GSMParams is a type for GSM parameters.
type GSMParams struct {
	ARFCN int32 `json:"Arfcn"`
	ULSC  int32 `json:"Ulsc"`
	BCC   int32 `json:"Bcc"`
}

// WCDMAParams is a type for WCDMA parameters.
type WCDMAParams struct {
	ARFCN      int32 `json:"Arfcn"`
	PSC        int32 `json:"Psc"`
	TargetCode int64 `json:"TargetCode"`
}

// LTEParams is a type for LTE parameters.
type LTEParams struct {
	ARFCN int32 `json:"Arfcn"`
	RNTI  int32 `json:"Rnti"`
	PCI   int32 `json:"Pci"`
}

// Point is a type for point data.
type Point struct {
	ID              int64   `json:"Id"`
	SettingsVersion int32   `json:"SettingsVersion"`
	Channel         int32   `json:"Channel"`
	SNR             float64 `json:"SNR"`
	RSSI            float64 `json:"RSSI"`
	TSf             float64 `json:"TSf"`
	TSi             int32   `json:"TSi"`
}

// GPSExt is a type for GPS data.
type GPSExt struct {
	Lat    float64 `json:"Lat"`
	Lon    float64 `json:"Lon"`
	Status int32   `json:"Status"`
}
