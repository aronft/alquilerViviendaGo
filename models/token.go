package models

//Token es el autentificador cuando el usuario inicia sesion
type Token struct {
	Token string `json:"token"`
}
