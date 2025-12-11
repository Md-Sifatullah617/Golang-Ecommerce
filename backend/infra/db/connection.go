package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "user=metalplus password=admin2025 host=localhost port =5432 dbname=ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbsource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbsource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
