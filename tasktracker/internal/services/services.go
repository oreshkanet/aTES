// Имплементация бизнес-логики приложения
package services

import "github.com/oreshkanet/aTES/tasktracker/internal/repository"

type Services struct {
	Users *UsersService
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Users: &UsersService{repos.Users},
	}
}
