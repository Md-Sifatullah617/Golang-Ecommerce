package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)
	mux := http.NewServeMux()
	wrappedMux := manager.WrappedMux(mux)
	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	println("🚀 Server is running at http://localhost" + addr)
	// globalRouter := middleware.Cors(mux)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
