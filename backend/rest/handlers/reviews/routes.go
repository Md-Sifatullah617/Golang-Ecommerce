package reviews

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(
	mux *http.ServeMux,
	manager *middleware.Manager,
) {
	mux.Handle("GET /reviews", manager.With(
		http.HandlerFunc(h.GetReviews),
	))
}
