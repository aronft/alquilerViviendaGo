package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claim sirve para vlaidar si la persona se autentico o no
type Claim struct {
	Propietario `json:"propietario"`
	jwt.StandardClaims
}
