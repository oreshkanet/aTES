package api

import (
	"fmt"
	"github.com/oreshkanet/aTES/auth/internal/domain"
)

type ErrorResponse struct {
	Error    bool   `json:"err"`
	ErrorMsg string `json:"err_msg"`
}

type SignUpResponse struct {
	User *domain.User `json:"user"`
}

type SignInResponse struct {
	Token string `json:"authorizer,omitempty"`
}

func newSignInResponse(token string) *SignInResponse {
	return &SignInResponse{
		token,
	}
}

func newSignUpResponse(user *domain.User) *SignUpResponse {
	return &SignUpResponse{
		user,
	}
}

func newErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		true, fmt.Sprintf("%s", err.Error()),
	}
}
