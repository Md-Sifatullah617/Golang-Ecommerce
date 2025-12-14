package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/users"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(u User) (*User, error)
}
