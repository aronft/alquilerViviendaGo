package models

//Departamento es la estrucura de la tabla departamentos
type Departamento struct {
	ID       uint     `json:"id"`
	Nombre   string   `json:"nombre"`
	Ciudades []Ciudad `json:"ciudad,omitempty"`
}
