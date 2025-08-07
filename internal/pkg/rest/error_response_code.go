package rest

type ApiError struct {
	Code       int
	CodeStatus string
}

func NewApiError(code int, codeStatus string) ApiError {
	return ApiError{
		Code:       code,
		CodeStatus: codeStatus,
	}
}

var (
	BadRequestError          = ApiError{BadRequestCode, BadRequestStatus}
	InternalServerErrorError = ApiError{InternalServerErrorCode, InternalServerErrorStatus}
	NotFoundError            = ApiError{NotFoundCode, NotFoundStatus}
	ConflictError            = ApiError{ConflictCode, ConflictStatus}

	UnauthorizedError = ApiError{UnauthorizedCode, UnauthorizedStatus}
	ForbiddenError    = ApiError{ForbiddenCode, ForbiddenStatus}

	ValidationError = ApiError{ValidationErrorCode, ValidationErrorStatus}
)

const (
	BadRequestCode          = 400
	InternalServerErrorCode = 500
	NotFoundCode            = 404
	ConflictCode            = 409

	UnauthorizedCode = 401
	ForbiddenCode    = 403

	ValidationErrorCode = 600
)

const (
	BadRequestStatus          = "Bad Request"
	InternalServerErrorStatus = "Internal Server Error"
	NotFoundStatus            = "Not Found"
	ConflictStatus            = "Conflict"

	UnauthorizedStatus = "Unauthorized"
	ForbiddenStatus    = "Forbidden"

	ValidationErrorStatus = "Validation Failed"
)
