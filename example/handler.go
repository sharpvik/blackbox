package main

import (
	"fmt"
	"net/http"

	"github.com/sharpvik/blackbox"
	"github.com/sharpvik/blackbox/filters"
	"github.com/sharpvik/blackbox/middleware"
)

func mainHandler(users []user) (rtr *blackbox.Router) {
	rtr = blackbox.New()

	// GET /api/{username}
	rtr.Subrouter().
		WithFilters(filters.Methods(http.MethodGet)).
		WithRoute(middleware.Params(
			middleware.ParamsList(
				middleware.ParamConst("api"),
				middleware.ParamString("username")),
			apiHandler(users)))

	// Every request not matched by the above subrouter, will be handled by
	// the default rtr.handler. It will set status 501 Not Implemented, and
	// write "Not Implemented" to the response body.
	return
}

func apiHandler(users []user) blackbox.HandlerFunc {
	return func(r *http.Request) *blackbox.Response {
		username := middleware.GetParams(r)["username"]
		fmt.Println("Request to view user info for user:", username)
		resp := blackbox.NewResponse()
		err := resp.EncodeJSON(findUser(users, username))
		if err != nil {
			panic("Failed to encode JSON for user: " + username)
		}
		return resp
	}
}
