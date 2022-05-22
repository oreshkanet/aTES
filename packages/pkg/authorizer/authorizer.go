package authorizer

type AuthToken interface {
	Generate(publicId string) (string, error)
	ParseToken(accessToken string, claims interface{}) error
}
