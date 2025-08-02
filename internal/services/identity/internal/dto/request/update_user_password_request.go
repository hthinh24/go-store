package request

import "github.com/hthinh24/go-store/services/identity/internal/errors"

type UpdateUserPasswordRequest struct {
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm"`
}

// Validate TODO - Implement validation logic for UpdateUserPasswordRequest
func (r *UpdateUserPasswordRequest) Validate() error {
	if r.NewPassword != r.NewPasswordConfirm {
		return errors.ErrPasswordMismatch{}
	}

	return nil
}
