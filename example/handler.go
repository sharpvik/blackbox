package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sharpvik/blackbox"
	"github.com/sharpvik/blackbox/filters"
)

func mainHandler(users []user) (rtr *blackbox.Router) {
	rtr = blackbox.New()

	rtr.Subrouter().
		WithFilters(filters.PathPrefix("/api/")).
		WithHandler(apiHandler(users))

	return
}

func apiHandler(users []user) blackbox.HandlerFunc {
	return func(r *http.Request) *blackbox.Response {
		username := strings.TrimPrefix(r.URL.String(), "/api/")
		fmt.Println("Request to view user info for user:", username)
		resp := blackbox.NewResponse().WithStatus(http.StatusOK)
		err := resp.EncodeJSON(findUser(users, username))
		if err != nil {
			panic("Failed to encode JSON for user: " + username)
		}
		return resp
	}
}
