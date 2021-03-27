package blackbox

import "net/http"

// Handler interface represents a functional handler that returns the Response
// instance taking net/http.Request as its only parameter.
type Handler interface {
	Handle(r *http.Request) *Response
}

// HandlerFunc is a special function type that implements the Handler interface.
type HandlerFunc func(r *http.Request) *Response

// Handle helps HandlerFunc implement the Handler interface.
func (h HandlerFunc) Handle(r *http.Request) *Response {
	return h(r)
}
