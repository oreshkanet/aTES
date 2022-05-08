package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/oreshkanet/aTES/auth/pkg/auth/models"
)

type Auth struct {
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuth(signingKey []byte, hashSalt string, expireDuration time.Duration) Auth {
	return Auth{
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

func (a *Auth) SignUp(ctx context.Context, user *models.User) error {
	user.Password = a.generatePasswordHash(user.Password)

	return nil // TODO: insert to repos
}

func (a *Auth) SignIn(ctx context.Context, user *models.User) (string, error) {
	user.Password = a.generatePasswordHash(user.Password)

	// user, err := // TODO: get from repos
	user.Role = "admin"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		UserName: user.Name,
		Role:     user.Role,
	})

	return token.SignedString(a.signingKey)
}

func (a *Auth) generatePasswordHash(password string) string {
	pwdHash := sha1.New()
	pwdHash.Write([]byte(password))
	pwdHash.Write([]byte(a.hashSalt))
	return fmt.Sprintf("%x", pwdHash.Sum(nil))
}
