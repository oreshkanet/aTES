package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/auth/internal/domain"
	"github.com/oreshkanet/aTES/auth/internal/service"
	"net/http"
)

type handler struct {
	auth service.AuthService
}

func NewHandler(aSvc service.AuthService) *handler {
	return &handler{
		auth: aSvc,
	}
}

func (h *handler) signUp(c *gin.Context) {
	req := new(SignUpRequest)
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
		return
	}

	user := &domain.User{
		PublicId: req.PublicId,
		Name:     req.Name,
		Password: req.Password,
		Role:     req.Password,
	}

	err := h.auth.SignUp(c.Request.Context(), user)
	if err != nil {
		if err == domain.ErrorUserAlreadyExists {
			c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, newErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, newSignUpResponse(user.PublicId))
}

func (h *handler) signIn(c *gin.Context) {
	req := new(SignInRequest)
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
		return
	}

	token, err := h.auth.SignIn(c.Request.Context(), req.PublicId, req.Password)
	if err != nil {
		if err == domain.ErrorInvalidAccessToken {
			c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
			return
		}

		if err == domain.ErrorUserDoesNotExist {
			c.AbortWithStatusJSON(http.StatusNotFound, newErrorResponse(err))
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, newErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, newSignInResponse(token))
}

func (h *handler) userChangeRole(c *gin.Context) {
	req := new(UserChangeRoleRequest)
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
		return
	}

	err := h.auth.ChangeRole(c.Request.Context(), req.PublicId, req.Role)
	if err != nil {
		if err == domain.ErrorUserAlreadyExists {
			c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, newErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, newUserChangeRoleResponse(req.PublicId))
}

func (h *handler) userUpdateProfile(c *gin.Context) {
	req := new(UserUpdateProfileRequest)
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
		return
	}

	user := &domain.User{
		PublicId: req.PublicId,
		Name:     req.Name,
		Password: req.Password,
	}

	err := h.auth.UpdateUserProfile(c.Request.Context(), user)
	if err != nil {
		if err == domain.ErrorUserAlreadyExists {
			c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, newErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, newUserUpdateProfileResponse(user.PublicId))
}
