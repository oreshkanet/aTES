package models

import "github.com/dgrijalva/jwt-go/v4"

type Claims struct {
	jwt.StandardClaims
	PublicId string `json:"public_id"`
}
