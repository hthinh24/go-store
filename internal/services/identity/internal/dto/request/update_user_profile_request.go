package request

import (
	"errors"
	"time"
)

type UpdateUserProfileRequest struct {
	Email       *string    `json:"email,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	FirstName   *string    `json:"first_name,omitempty"`
	Avatar      *string    `json:"avatar,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
}

// Validate TODO - Add validation for email format, phone number format, etc.
func (r *UpdateUserProfileRequest) Validate() error {
	// Check if at least one field is provided
	if r.Email == nil && r.LastName == nil && r.FirstName == nil &&
		r.Avatar == nil && r.Gender == nil && r.PhoneNumber == nil &&
		r.DateOfBirth == nil {
		return errors.New("at least one field must be provided for update")
	}

	return nil
}
