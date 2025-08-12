package request

type UpdateItemQuantityRequest struct {
	ItemID   int64 `json:"item_id" binding:"required"`
	Quantity int   `json:"quantity" binding:"required,min=1"`
}
