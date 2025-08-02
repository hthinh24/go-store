package rest

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewAPIResponse(code int, message string, data interface{}) *APIResponse {
	return &APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
