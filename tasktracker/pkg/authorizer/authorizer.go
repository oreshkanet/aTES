// Package authorizer - пакет для организации jwt-авторизации в системе
package authorizer

type AuthToken interface {
	Generate(publicId string) (string, error)
	ParseToken(accessToken string) (string, error)
}
