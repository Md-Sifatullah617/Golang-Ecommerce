package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/products"
	"ecommerce/rest/handlers/reviews"
	"ecommerce/rest/handlers/users"
	"ecommerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	middlewares:=middleware.NewMiddlewares(cnf)

	productHandler := products.NewHandler(middlewares)
	userHandler := users.NewHandler()
	reviewHandler := reviews.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()
}
