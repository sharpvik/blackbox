package filters

import (
	"net/http"
	"strings"

	"github.com/sharpvik/blackbox"
)

// Path filter filters out requests with URIs that exactly matche p.
func Path(p string) blackbox.FilterFunc {
	return func(r *http.Request) bool {
		return r.URL.String() == p
	}
}

// PathPrefix filters out requests with URIs that are prefixed with p.
func PathPrefix(p string) blackbox.FilterFunc {
	return func(r *http.Request) bool {
		return strings.HasPrefix(r.URL.String(), p)
	}
}
