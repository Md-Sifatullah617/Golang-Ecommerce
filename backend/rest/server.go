package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/products"
	"ecommerce/rest/handlers/users"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *products.Handler
	userHandler    *users.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *products.Handler,
	userHandler *users.Handler,

) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	mux := http.NewServeMux()
	wrappedMux := manager.WrappedMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	println("🚀 Server is running at http://localhost" + addr)
	// globalRouter := middleware.Cors(mux)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
