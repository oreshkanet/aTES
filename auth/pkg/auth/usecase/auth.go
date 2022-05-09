package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/oreshkanet/aTES/auth/pkg/auth/models"
	"github.com/oreshkanet/aTES/auth/pkg/auth/repository"
	"github.com/oreshkanet/aTES/auth/pkg/auth/transport"
)

type Auth struct {
	repos          *repository.UserRepository
	topics         *transport.Transports
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuth(repos *repository.UserRepository, topics *transport.Transports, signingKey []byte, hashSalt string, expireDuration time.Duration) Auth {
	return Auth{
		repos:          repos,
		topics:         topics,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

func (a *Auth) SignUp(ctx context.Context, user *models.User) error {
	// Хэшируем пароль нового пользователя
	user.Password = a.generatePasswordHash(user.Password)

	// Добавляем пользователя в БД
	if err := a.repos.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("Create user:%w", err)
	}

	// Отправляем сообщение в топик
	if err := a.topics.PubUser(ctx, user); err != nil {
		return fmt.Errorf("Create user:%w", err)
	}

	return nil
}

func (a *Auth) SignIn(ctx context.Context, user *models.User) (string, error) {
	// Ищем в БД пользователя по имени
	userDB, err := a.repos.SelectUserByName(ctx, user.Name)
	if err != nil {
		return "", err
	}

	// Проверяем совпадают ли пароли
	user.Password = a.generatePasswordHash(user.Password)
	if user.Password != userDB.Password {
		return "", fmt.Errorf("incorrect password")
	}

	// Генерируем токен доступа JWT
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
