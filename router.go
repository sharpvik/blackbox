package blackbox

import "net/http"

// Router is the main feature of the blackbox package. It allows you to build
// extensible, testable and typesafe routing for your web app or microservice.
//
// Router implements
//   - net/http.Handler
//   - Filter
//   - Handler
//   - Route
//
type Router struct {
	filters []Filter
	routes  []Route
	handler Handler
}

// New returns a newly create Router instance with pre-allocated filters and
// routes and a default catch-all handler that always responds with status
// 501 Not Implemented. Feel free to customise it using available methods!
func New() *Router {
	return &Router{
		filters: make([]Filter, 0),
		routes:  make([]Route, 0),
		handler: statusAndMessage(http.StatusNotFound, "Not Found"),
	}
}

// ServeHTTP helps Router implement net/http.Handler interface for effective
// interop with the net/http.Server.
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rtr.Handle(r).Respond(w)
}

// Accepts helps Router implement the Filter and Route interfaces.
func (rtr *Router) Accepts(r *http.Request) bool {
	for _, filter := range rtr.filters {
		if !filter.Accepts(r) {
			return false
		}
	}
	return true
}

// Handle helps Router implement the Handler and Route interfaces.
func (rtr *Router) Handle(r *http.Request) (resp *Response) {
	resp = rtr.propagate(r)
	if resp != nil {
		return
	}
	return rtr.handler.Handle(r)
}

// WithFilters accepts multiple Filter's and adds them to the Router, returning
// a reference to the same Router instance.
func (rtr *Router) WithFilters(filters ...Filter) *Router {
	rtr.filters = append(rtr.filters, filters...)
	return rtr
}

// WithFilterFuncs accepts mutilple FilterFunc's and adds them to the Router,
// returning a reference to the same Router instance.
func (rtr *Router) WithFilterFuncs(filters ...FilterFunc) *Router {
	for _, filter := range filters {
		rtr.filters = append(rtr.filters, filter)
	}
	return rtr
}

// WithRoute accepts a Route, adds it to the registered routes, and returns
// reference to the accepting Router.
func (rtr *Router) WithRoute(sub Route) *Router {
	rtr.routes = append(rtr.routes, sub)
	return rtr
}

// Subrouter adds a newly created subrouter instance to the Router, returning
// a reference to that subrouter for future modification.
func (rtr *Router) Subrouter() (sub *Router) {
	sub = New()
	rtr.routes = append(rtr.routes, sub)
	return
}

// WithHandler accepts a Handler and adds it to the Router, returning a
// reference to the same Router instance.
func (rtr *Router) WithHandler(handler Handler) *Router {
	rtr.handler = handler
	return rtr
}

// WithHandlerFunc accepts a HandlerFunc and adds it to the Router, returning
// a reference to the same Router instance.
func (rtr *Router) WithHandlerFunc(f HandlerFunc) *Router {
	rtr.handler = f
	return rtr
}

func (rtr *Router) propagate(r *http.Request) *Response {
	for _, route := range rtr.routes {
		if route.Accepts(r) {
			return route.Handle(r)
		}
	}
	return nil
}
