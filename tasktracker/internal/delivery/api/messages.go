package api

type TaskAddRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskAddResponse struct {
	PublicId string `json:"public_id"`
}

type ErrorResponse struct {
	Error    bool   `json:"err"`
	ErrorMsg string `json:"err_msg"`
}
