package users

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type Handler struct {
	cnf      *config.Config
	userRepo repo.UserRepo
}

func NewHandler(
	userRepo repo.UserRepo,
	cnf *config.Config,
) *Handler {
	return &Handler{
		cnf:      cnf,
		userRepo: userRepo,
	}
}
