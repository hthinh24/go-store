package service

import (
	"github.com/hthinh24/go-store/services/identity"
	"log"
)

type userService struct {
	Logger     *log.Logger
	Repository *identity.UserRepository
}

func NewUserService(logger *log.Logger, repository *identity.UserRepository) *userService {
	return &userService{
		Logger:     logger,
		Repository: repository,
	}
}
