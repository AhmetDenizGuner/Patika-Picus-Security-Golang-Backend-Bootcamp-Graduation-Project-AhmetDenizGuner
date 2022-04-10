package shared

type ApiOkResponse struct {
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type ApiErrorResponse struct {
	IsSuccess    bool   `json:"is_success"`
	ErrorMessage string `json:"error_message"`
}
