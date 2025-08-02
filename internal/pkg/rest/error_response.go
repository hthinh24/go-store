package rest

type ErrorResponse struct {
	ApiError
	Message string `json:"message"`
}

func NewErrorResponse(apiError ApiError, message string) ErrorResponse {
	return ErrorResponse{
		ApiError: apiError,
		Message:  message,
	}
}
