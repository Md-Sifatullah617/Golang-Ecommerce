package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	prdHandler "ecommerce/rest/handlers/products"
	usrHandler "ecommerce/rest/handlers/users"
	"ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	// Load config
	cnf := config.GetConfig()

	// Connect to database
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println("failed to connect database:", err)
		os.Exit(1)
	}
	defer dbCon.Close()

	// Run migrations
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println("Migration failed:", err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	// Repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// Domains
	prdcSvc := product.NewService(productRepo)
	usrSvc := user.NewService(userRepo)

	productHandler := prdHandler.NewHandler(middlewares, prdcSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
