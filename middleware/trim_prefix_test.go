package middleware

import (
	"net/http"
	"testing"

	"github.com/sharpvik/blackbox"
	"github.com/sharpvik/blackbox/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestTrimPrefix(t *testing.T) {
	api := TrimPrefix("/api", blackbox.HandlerFunc(respondWithURI))
	req := test_utils.Get(t, "/api/public")
	assert.True(t, api.Accepts(req))
	assert.Equal(t, "/public", api.Handle(req).Body.String())
}

func respondWithURI(r *http.Request) *blackbox.Response {
	return blackbox.NewResponse().
		WriteString(r.URL.String()).
		WithStatus(http.StatusOK)
}
