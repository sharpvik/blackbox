package blackbox

import (
	"net/http"
	"net/http/httptest"
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
	assert.Equal(t, resp.Status, http.StatusNotImplemented)
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

func TestWithCookie(t *testing.T) {
	rtr := blackbox.New().WithHandler(
		test_utils.RespondWithStatusAndCookie(http.StatusOK, "cookie", "ok"))
	rec := httptest.NewRecorder()
	rtr.ServeHTTP(rec, test_utils.Get(t, "/"))
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	cookies := rec.Result().Cookies()
	assert.Equal(t, 1, len(cookies))
	cookie := cookies[0]
	assert.Equal(t, "cookie", cookie.Name)
	assert.Equal(t, "ok", cookie.Value)
}
