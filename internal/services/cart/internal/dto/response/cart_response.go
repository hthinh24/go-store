package response

type CartResponse struct {
	ID     int64               `json:"id"`
	UserID int64               `json:"user_id"`
	Status string              `json:"status"`
	Items  *[]CartItemResponse `json:"items"`
}
