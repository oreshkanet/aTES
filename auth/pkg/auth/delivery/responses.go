package delivery

import (
	"github.com/oreshkanet/aTES/auth/pkg/auth/models"
)

type ErrorResponse struct {
	Error    bool   `json:"err"`
	ErrorMsg string `json:"err_msg"`
}

type SignUpResponse struct {
	user *models.User `json:"user"`
}

type SignInResponse struct {
	Token string `json:"token,omitempty"`
}

func newSignInResponse(token string) *SignInResponse {
	return &SignInResponse{
		token,
	}
}

func newSignUpResponse(user *models.User) *SignUpResponse {
	return &SignUpResponse{
		user,
	}
}

func newErrorResponse(errorMessage string) *ErrorResponse {
	return &ErrorResponse{
		true, errorMessage,
	}
}
