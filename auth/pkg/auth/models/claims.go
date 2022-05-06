package models

import "github.com/dgrijalva/jwt-go/v4"

type Claims struct {
	jwt.StandardClaims
	UserName string `json:"username"`
	Role     string `json:"role"`
}
