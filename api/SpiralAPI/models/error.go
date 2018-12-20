package models

type errorResponse struct {
	ErrorCode int64 `json:"error_code"`
	Message string `json:"message"`
}
