package repo

import "fmt"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Find(email, password string) (*User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(u User) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(u User) (*User, error) {
	if u.ID != 0 {
		return &u, nil
	}

	u.ID = len(r.users) + 1
	r.users = append(r.users, u)
	return &u, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	for i := range r.users {
		if r.users[i].Email == email && r.users[i].Password == password {
			return &r.users[i], nil
		}
	}
	return nil, fmt.Errorf("user with email ID %s not found", email)
}
