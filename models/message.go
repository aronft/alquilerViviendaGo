package models

//Message es para enviar mensajes y los codigos el cliente
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
