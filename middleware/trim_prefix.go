package middleware

import (
	"net/http"
	"strings"

	"github.com/sharpvik/blackbox"
	"github.com/sharpvik/blackbox/filters"
)

// TrimPrefix constructs TrimPrefixRoute by accepting a prefix to be matched
// against, and a blackbox.Handler.
func TrimPrefix(p string, h blackbox.Handler) *Middleware {
	return NewMiddleware(
		filters.PathPrefix(p),
		blackbox.HandlerFunc(func(r *http.Request) *blackbox.Response {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, p)
			return h.Handle(r)
		}),
	)
}
