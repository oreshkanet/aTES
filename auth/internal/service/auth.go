package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/oreshkanet/aTES/auth/internal/domain"
	"github.com/oreshkanet/aTES/auth/internal/events"
	"github.com/oreshkanet/aTES/auth/internal/repository"
	"github.com/oreshkanet/aTES/packages/pkg/authorizer"
)

type auth struct {
	repos     repository.UserRepository
	events    events.Producer
	authToken authorizer.AuthToken
	hashSalt  string
}

func newAuth(
	repos repository.UserRepository,
	events events.Producer,
	authToken authorizer.AuthToken,
	hashSalt string,
) *auth {
	return &auth{
		repos:     repos,
		events:    events,
		hashSalt:  hashSalt,
		authToken: authToken,
	}
}

func (a *auth) SignUp(ctx context.Context, user *domain.User) error {
	// Хэшируем пароль нового пользователя
	user.Password = a.generatePasswordHash(user.Password)

	// Добавляем пользователя в БД
	if err := a.repos.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("Create user:%w", err)
	}

	// Отправляем событие о создании нового пользователя в stream
	if err := a.events.UserCreated(ctx, user); err != nil {
		return fmt.Errorf("Create user:%w", err)
	}

	return nil
}

func (a *auth) SignIn(ctx context.Context, publicId string, pwd string) (string, error) {
	// Ищем в БД пользователя по имени
	userDB, err := a.repos.FindUserByPublicId(ctx, publicId)
	if err != nil {
		return "", err
	}

	// Проверяем совпадают ли пароли
	pwdHash := a.generatePasswordHash(pwd)
	if pwdHash != userDB.Password {
		return "", fmt.Errorf("incorrect password")
	}

	// Генерируем токен доступа JWT
	token, err := a.authToken.Generate(userDB.PublicId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *auth) generatePasswordHash(password string) string {
	pwdHash := sha1.New()
	pwdHash.Write([]byte(password))
	pwdHash.Write([]byte(a.hashSalt))
	return fmt.Sprintf("%x", pwdHash.Sum(nil))
}

func (a *auth) ChangeRole(ctx context.Context, publicId string, role string) error {
	// Ищем в БД пользователя по имени
	userDB, err := a.repos.FindUserByPublicId(ctx, publicId)
	if err != nil {
		return err
	}

	// Меняем роль пользователя
	userDB.Role = role

	// Апдейтим в БД
	err = a.repos.UpdateUser(ctx, userDB)
	if err != nil {
		return err
	}

	// Публикуем событие изменения роли пользователя
	err = a.events.UserRoleChanged(ctx, userDB)
	if err != nil {
		return err
	}

	return nil
}

func (a *auth) UpdateUserProfile(ctx context.Context, user *domain.User) error {
	// Ищем в БД пользователя по имени
	userDB, err := a.repos.FindUserByPublicId(ctx, user.PublicId)
	if err != nil {
		return err
	}

	// Обновляем данные профиля пользователя
	userDB.Name = user.Name

	// Обновляем данные в БД
	err = a.repos.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	// Стримим событие изменения пользователя
	err = a.events.UserUpdated(ctx, userDB)
	if err != nil {
		return err
	}

	return nil
}
