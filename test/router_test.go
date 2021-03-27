package blackbox

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sharpvik/blackbox"
	"github.com/sharpvik/blackbox/filters"
)

func TestDefaultRouter(t *testing.T) {
	rtr := blackbox.New()
	req := get(t, "/")
	resp := rtr.Handle(req)
	assert.Equal(t, resp.Status, http.StatusNotImplemented)
}

func TestWithSubroutesFiltersAndHandlers(t *testing.T) {
	rtr := blackbox.New().
		WithHandler(
			respondWithStatusAndMessage(
				http.StatusOK, "default"))

	rtr.Subrouter().
		WithFilters(filters.Path("/api")).
		WithHandler(
			respondWithStatusAndMessage(
				http.StatusOK, "API"))

	rtr.Subrouter().
		WithFilters(filters.Path("/pub")).
		WithHandler(
			respondWithStatusAndMessage(
				http.StatusOK, "public"))

	getAndCheckStatusAndBody(t, rtr, "/api", http.StatusOK, "API")
	getAndCheckStatusAndBody(t, rtr, "/pub", http.StatusOK, "public")
	getAndCheckStatusAndBody(t, rtr, "/", http.StatusOK, "default")
}

func getAndCheckStatusAndBody(
	t *testing.T, rtr *blackbox.Router, uri string, status int, body string) {
	req := get(t, uri)
	resp := rtr.Handle(req)
	assert.Equal(t, status, resp.Status, "unexpected response status")
	assert.Equal(t, body, resp.Body.String(), "unexpected response body")
}

func get(t *testing.T, uri string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	assert.NoError(t, err, "unexpected error creating request")
	return req
}

func respondWithStatusAndMessage(
	status int, message string) blackbox.HandlerFunc {
	return func(*http.Request) *blackbox.Response {
		resp := blackbox.NewResponse().WithStatus(status)
		resp.WriteString(message)
		return resp
	}
}
