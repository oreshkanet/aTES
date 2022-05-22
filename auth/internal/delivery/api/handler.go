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
	user := new(domain.User)
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
		return
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

	c.JSON(http.StatusOK, newSignUpResponse(user))
}

func (h *handler) signIn(c *gin.Context) {
	user := new(domain.User)
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(err))
		return
	}

	token, err := h.auth.SignIn(c.Request.Context(), user)
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
