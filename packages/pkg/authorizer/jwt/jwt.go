package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type Token struct {
	signingKey     string
	expireDuration time.Duration
	signingMethod  *jwt.SigningMethodHMAC
}

func NewJwtToken(signingKey string, expireDuration time.Duration) *Token {
	return &Token{
		signingKey:     signingKey,
		expireDuration: expireDuration,
		signingMethod:  jwt.SigningMethodHS256,
	}
}

func (a *Token) Generate(publicId string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
				IssuedAt:  jwt.At(time.Now()),
			},
			PublicId: publicId,
		},
	)
	return token.SignedString(a.signingKey)
}

func (a *Token) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return "", nil
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.PublicId, nil
	}

	return "", fmt.Errorf("invalid access token")
}
