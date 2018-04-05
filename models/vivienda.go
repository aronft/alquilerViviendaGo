package models

//Vivienda es la estructura de la tabla viviendas
type Vivienda struct {
	ID            uint   `json:"id"`
	PropietarioID uint   `json:"propietarioId"`
	CiudadID      uint   `json:"ciudadId"`
	Habitantes    uint   `json:"habitantes"`
	DirCalle      string `json:"dirCalle"`
	DirNumero     uint   `json:"dirNumero"`
	Descripcion   string `json:"descripcion"`
}
