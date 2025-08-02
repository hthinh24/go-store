package request

import (
	"github.com/hthinh24/go-store/services/identity/internal/constants"
	"time"
)

// CreateUserRequest TODO - Add validation tags to the struct fields
type CreateUserRequest struct {
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	ProviderID   string    `json:"provider_id"`
	ProviderName string    `json:"provider_name"`
	LastName     string    `json:"last_name"`
	FirstName    string    `json:"first_name"`
	Avatar       string    `json:"avatar"`
	Gender       string    `json:"gender"`
	PhoneNumber  string    `json:"phone_number"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Status       string    `json:"status"`
}

func (c *CreateUserRequest) Validate() error {
	c.Gender = constants.GetGender(c.Gender)
	c.Status = constants.GetStatus(c.Status)

	return nil
}
