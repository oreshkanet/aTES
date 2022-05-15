package api

type GetBalanceRequest struct {
	UserPublicId string `json:"user_public_id"`
}

type GetBalanceResponse struct {
	Balance float32 `json:"balance"`
}

type ErrorResponse struct {
	Error    bool   `json:"err"`
	ErrorMsg string `json:"err_msg"`
}
