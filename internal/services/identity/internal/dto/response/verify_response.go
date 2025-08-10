package response

type VerifyResponse struct {
	UserID      string   `json:"user_id"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}
