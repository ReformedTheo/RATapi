package models

type RequestData struct {
	HEX      string `json:"hex"`      //hex refering to the RACK
	Coil_HEX string `json:"coil_hex"` //hex refering to the Coils, only usable if going to maintence
	Client   string `json:"client"`
	Coils    []Coil `json:"coils"`
	Status   int    `json:"status"`
}

type Coil struct {
	HEX    string `json:"hex"`
	Status int    `json:"status"`
	Cycles int    `json:"cycles"`
	Date   string `json:"date"`
}

type Rack struct {
	Coils  []Coil `json:"coils"`
	Client string `json:"client"`
	Date   string `json:"date"`
	Status string `json:"status"`
}
