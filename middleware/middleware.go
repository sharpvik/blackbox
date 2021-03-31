package middleware

import (
	"net/http"

	"github.com/sharpvik/blackbox"
)

type Middleware struct {
	filter  blackbox.Filter
	handler blackbox.Handler
}

func NewMiddleware(f blackbox.Filter, h blackbox.Handler) *Middleware {
	return &Middleware{
		filter:  f,
		handler: h,
	}
}

func (m *Middleware) Accepts(r *http.Request) bool {
	return m.filter.Accepts(r)
}

func (m *Middleware) Handle(r *http.Request) *blackbox.Response {
	return m.handler.Handle(r)
}
