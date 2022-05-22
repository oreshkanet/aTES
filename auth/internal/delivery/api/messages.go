package api

import (
	"fmt"
	"github.com/oreshkanet/aTES/auth/internal/domain"
)

type ErrorResponse struct {
	Error    bool   `json:"err"`
	ErrorMsg string `json:"err_msg"`
}

func newErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		true, fmt.Sprintf("%s", err.Error()),
	}
}

type SignUpRequest struct {
	*domain.User
}

type SignUpResponse struct {
	PublicId string `json:"public_id"`
}

func newSignUpResponse(publicId string) *SignUpResponse {
	return &SignUpResponse{
		publicId,
	}
}

type SignInRequest struct {
	PublicId string `json:"public_id"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func newSignInResponse(token string) *SignInResponse {
	return &SignInResponse{token}
}

type UserUpdateProfileRequest struct {
	*domain.User
}

type UserUpdateProfileResponse struct {
	PublicId string `json:"public_id"`
}

func newUserUpdateProfileResponse(publicId string) *UserUpdateProfileResponse {
	return &UserUpdateProfileResponse{publicId}
}

type UserChangeRoleRequest struct {
	PublicId string `json:"public_id"`
	Role     string `json:"role"`
}

type UserChangeRoleResponse struct {
	PublicId string `json:"public_id"`
}

func newUserChangeRoleResponse(publicId string) *UserChangeRoleResponse {
	return &UserChangeRoleResponse{publicId}
}
