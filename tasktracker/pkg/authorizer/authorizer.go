// Package authorizer - пакет для организации jwt-авторизации в системе
//TODO: перевезти в отдельный пакет packages
package authorizer

type AuthToken interface {
	Generate(publicId string) (string, error)
	ParseToken(accessToken string) (string, error)
}
