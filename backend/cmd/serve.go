package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/products"
	"ecommerce/rest/handlers/users"
	"ecommerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	middlewares := middleware.NewMiddlewares(cnf)
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	productHandler := products.NewHandler(middlewares, productRepo)
	userHandler := users.NewHandler(userRepo, cnf)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
