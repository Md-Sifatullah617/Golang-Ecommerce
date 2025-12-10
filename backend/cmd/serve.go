package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/products"
	"ecommerce/rest/handlers/reviews"
	"ecommerce/rest/handlers/users"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := products.NewHandler()
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
