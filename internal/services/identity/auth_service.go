package identity

import (
	"github.com/hthinh24/go-store/services/identity/internal/dto/request"
	"github.com/hthinh24/go-store/services/identity/internal/dto/response"
)

type AuthService interface {
	Login(request request.AuthRequest) (response.AuthResponse, error)
	Logout(token string) error
}
