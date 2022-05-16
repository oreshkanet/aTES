package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/analytics/internal/domain"
	"github.com/oreshkanet/aTES/analytics/internal/service"
	"net/http"
)

type Handler struct {
	analytic service.AnalyticService
}

func NewHandler(analyticSvc service.AnalyticService) *Handler {
	return &Handler{analytic: analyticSvc}
}

func (h *Handler) getNegativeBalance(c *gin.Context) {
	var err error
	var users []*domain.User
	users, err = h.analytic.GetNegativeBalance(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorResponse{
			true, fmt.Sprintf("%s", err.Error()),
		})
	}

	c.JSON(http.StatusOK, &GetNegativeBalanceResponse{
		Rows: users,
	})
}
