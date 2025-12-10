package products

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
			middleware.AuthenticateJwt,
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			middleware.AuthenticateJwt,
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProductById),
			middleware.AuthenticateJwt,
		),
	)

	mux.Handle(
		"PATCH /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			middleware.AuthenticateJwt,
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			middleware.AuthenticateJwt,
		),
	)
}
