package filters

import (
	"net/http"

	"github.com/sharpvik/blackbox"
)

// Methods filter allows us to filter requests by method. This function accepts
// multiple filters as its parameters in case your subroute works with multiple
// different HTTP methods.
func Methods(methods ...string) blackbox.FilterFunc {
	return func(r *http.Request) bool {
		return stringSliceHas(methods, r.Method)
	}
}

func stringSliceHas(s []string, elem string) bool {
	for _, each := range s {
		if each == elem {
			return true
		}
	}
	return false
}
