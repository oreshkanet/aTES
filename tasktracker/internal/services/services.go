// Имплементация бизнес-логики приложения
package services

type Services struct {
	Users *UsersService
}

func NewServices() *Services {
	return &Services{
		Users: &UsersService{},
	}
}
