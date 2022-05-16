package api

import "github.com/oreshkanet/aTES/analytics/internal/domain"

type GetNegativeBalanceResponse struct {
	Rows []*domain.User `json:"rows"`
}

type ErrorResponse struct {
	Error    bool   `json:"err"`
	ErrorMsg string `json:"err_msg"`
}
