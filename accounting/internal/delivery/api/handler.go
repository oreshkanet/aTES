package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/accounting/internal/service"
	"net/http"
)

type Handler struct {
	accService service.AccountService
}

func NewHandler(service service.AccountService) *Handler {
	return &Handler{accService: service}
}

func (h *Handler) getBalance(c *gin.Context) {
	var err error
	req := new(GetBalanceRequest)
	err = c.BindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &ErrorResponse{
			true, fmt.Sprintf("%s", err.Error()),
		})
		return
	}

	var balance float32
	balance, err = h.accService.GetBalance(c.Request.Context(), req.UserPublicId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorResponse{
			true, fmt.Sprintf("%s", err.Error()),
		})
	}

	c.JSON(http.StatusOK, &GetBalanceResponse{
		Balance: balance,
	})
}
