package blackbox

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sharpvik/blackbox"
	"github.com/sharpvik/blackbox/filters"
	"github.com/sharpvik/blackbox/test_utils"
)

func TestDefaultRouter(t *testing.T) {
	rtr := blackbox.New()
	req := test_utils.Get(t, "/")
	resp := rtr.Handle(req)
	assert.Equal(t, resp.Status, http.StatusNotFound)
}

func TestWithSubroutesFiltersAndHandlers(t *testing.T) {
	rtr := blackbox.New().
		WithHandler(
			test_utils.RespondWithStatusAndMessage(
				http.StatusOK, "default"))

	rtr.Subrouter().
		WithFilters(filters.Path("/api")).
		WithHandler(
			test_utils.RespondWithStatusAndMessage(
				http.StatusOK, "API"))

	rtr.Subrouter().
		WithFilters(filters.Path("/pub")).
		WithHandler(
			test_utils.RespondWithStatusAndMessage(
				http.StatusOK, "public"))

	test_utils.GetAndCheckStatusAndBody(t, rtr, "/api", http.StatusOK, "API")
	test_utils.GetAndCheckStatusAndBody(t, rtr, "/pub", http.StatusOK, "public")
	test_utils.GetAndCheckStatusAndBody(t, rtr, "/", http.StatusOK, "default")
}
