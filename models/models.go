package models

// RequestData define a estrutura para os dados de entrada da sua requisição.
type RequestData struct {
	HEX    string `json:"hex"`
	State  int    `json:"state"`
	Client string `json:"client"`
}
type Coil struct {
	HEX    string
	Status int
	Cycles int
	Client string
}
