package model

/*
{
  "UnitID": "unit1",
  "Config": {
    "Network": 0,
    "GSMParams": {
      "Arfcn": 540,
      "Ulsc": 65,
      "Bcc": 2
    },
    "WCDMAParams": {
      "Arfcn": 10788,
      "TargetCode": 8758542,
      "Psc": 369
    },
    "LTEParams": {
      "Arfcn": 1400,
      "Rnti": 5,
      "Pci": 0
    },
    "Slot": "0",
    "Version": 2
  },
  "Points": [
    {
      "Id": 552,
      "SettingsVersion": 0,
      "Channel": 65,
      "SNR": 33.2,
      "RSSI": -28.3,
      "TSi": 139,
      "TSf": 0.351451076923
    }
  ],
  "GPSExt": {
    "Status": 0,
    "Lat": 32.097839,
    "Lon": 34.846317
  }
}
*/

// Request is a type for http request.
type Request struct {
	UnitID string  `json:"UnitId"`
	Config Config  `json:"Config"`
	Points []Point `json:"Points"`
	GpsExt GPSExt  `json:"GpsExt"`
}
