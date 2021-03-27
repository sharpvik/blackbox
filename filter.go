package blackbox

import "net/http"

// Filter represents any filtering entity that can decide whether it wants to
// deal with a given net/http.Request through its Accept predicate.
type Filter interface {
	Accepts(r *http.Request) bool
}

// FilterFunc is a special function type that implements the Filter interface.
type FilterFunc func(*http.Request) bool

// Accepts helps FilterFunc implement the Filter interface.
func (f FilterFunc) Accepts(r *http.Request) bool {
	return f(r)
}
