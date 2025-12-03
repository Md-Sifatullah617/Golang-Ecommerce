package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	handler := next

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

func (mngr *Manager) WrappedMux(next http.Handler) http.Handler {
	handler := http.Handler(next)
	for i := range mngr.globalMiddlewares {
		handler = mngr.globalMiddlewares[i](handler)
	}
	return handler
}
