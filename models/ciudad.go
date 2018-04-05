package models

//Ciudad es la esstrucutra de la tabla ciudades
type Ciudad struct {
	ID             uint       `json:"id"`
	DepartamentoID string     `json:"departamentoId"`
	Nombre         string     `json:"nombre"`
	Viviendas      []Vivienda `json:"viviendas,omitempty"`
}
