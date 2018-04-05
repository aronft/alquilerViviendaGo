package models

//Propietario es la estructura para el propietario
type Propietario struct {
	ID              uint       `json:"id"`
	NomNombres      string     `json:"nomNombres"`
	NomApellidos    string     `json:"nomApellidos"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Password        string     `json:"password,omitempty"`
	ConfirmPassword string     `json:"confirmPassword, omitempty"`
	Picture         string     `json:"picture"`
	Viviendas       []Vivienda `json:"viviendas,omitempty`
}
