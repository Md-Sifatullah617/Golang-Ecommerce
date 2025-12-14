package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/user"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r *userRepo) Create(u domain.User) (*domain.User, error) {
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

func (r *userRepo) Find(email, password string) (*domain.User, error) {
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
	var user domain.User
	err := r.dbCon.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid email & password")
		}
		return nil, fmt.Errorf("database error %w", err)
	}
	return &user, nil
}
