package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Find(email, password string) (*User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(u User) (*User, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r *userRepo) Create(u User) (*User, error) {
	query := `
	INSERT INTO users (
		first_name,
		last_name,
		email,
		password,
		is_shop_owner
	)
	VALUES (
		:first_name,
		:last_name,
		:email,
		:password,
		:is_shop_owner
	)
	RETURNING id
	`

	var userID int
	rows, err := r.dbCon.NamedQuery(query, u)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	defer rows.Close()
	if rows.Next() {

		if err := rows.Scan(&userID); err != nil {
			return nil, fmt.Errorf("failed to scan user ID: %w", err)
		}
	}
	u.ID = userID
	return &u, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	query := `
        SELECT 
            id,
            first_name,
            last_name,
            email,
            password,
            is_shop_owner
        FROM users
        WHERE email = $1 AND password = $2
        LIMIT 1
    `
	var user User
	err := r.dbCon.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid email & password")
		}
		return nil, fmt.Errorf("database error %w", err)
	}
	return &user, nil
}
