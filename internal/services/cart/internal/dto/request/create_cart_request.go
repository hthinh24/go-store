package request

type CreateCartRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}
