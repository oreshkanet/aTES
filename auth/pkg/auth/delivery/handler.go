package delivery

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/auth/pkg/auth/models"
	"github.com/oreshkanet/aTES/auth/pkg/auth/usecase"
)

type handler struct {
	auth usecase.Auth
}

func newHandler(auth usecase.Auth) *handler {
	return &handler{
		auth: auth,
	}
}

func RegisterServerEndpoits(router *gin.RouterGroup, auth usecase.Auth) {
	h := newHandler(auth)

	router.POST("/sign_up", h.signUp)
	router.POST("/sign_in", h.signIn)
}

func (h *handler) signUp(c *gin.Context) {
	user := new(models.User)
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(
			fmt.Sprintf("%s", err.Error()),
		))
		return
	}

	if err := h.auth.SignUp(c.Request.Context(), user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newSignUpResponse(user))
}

func (h *handler) signIn(c *gin.Context) {
	user := new(models.User)
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(
			fmt.Sprintf("%s", err.Error()),
		))
		return
	}

	token, err := h.auth.SignIn(c.Request.Context(), user)
	if err != nil {
		if err == models.ErrorInvalidAccessToken {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				newErrorResponse(fmt.Sprintf("%s", err.Error())))
			return
		}

		if err == models.ErrorUserDoesNotExist {
			c.AbortWithStatusJSON(http.StatusNotFound,
				newErrorResponse(fmt.Sprintf("%s", err.Error())))
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, newErrorResponse(
			fmt.Sprintf("%s", err.Error()),
		))
		return
	}

	c.JSON(http.StatusOK, newSignInResponse(token))
}
