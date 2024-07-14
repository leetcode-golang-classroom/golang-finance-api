package response

import "time"

type ApiError struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data"`
}

func (e *ApiError) Error() string {
	return e.Message
}

type ApiMetaData struct {
	Timestamp time.Time `json:"timestamp"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
}

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Meta    ApiMetaData `json:"meta"`
	Error   *ApiError   `json:"error"`
}
