package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/products"
	"ecommerce/rest/handlers/users"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbCon)

	productHandler := products.NewHandler(middlewares, productRepo)
	userHandler := users.NewHandler(userRepo, cnf)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
