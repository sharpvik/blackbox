package middleware

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sharpvik/blackbox"
	"github.com/sharpvik/blackbox/test_utils"
)

func TestParams(t *testing.T) {
	api := Params(
		ParamsList(
			ParamConst("api"),
			ParamString("handle"),
			ParamInt("id")),
		blackbox.HandlerFunc(checkParams))
	assert.False(t, api.Accepts(test_utils.Get(t, "/api/sharpvik")))
	assert.False(t, api.Accepts(test_utils.Get(t, "/api/sharpvik/42/lol")))
	req := test_utils.Get(t, "/api/sharpvik/42")
	assert.True(t, api.Accepts(req))
	assert.Equal(t, http.StatusOK, api.Handle(req).Status)
}

func checkParams(r *http.Request) *blackbox.Response {
	params := GetParams(r)
	if ok := params["handle"] == "sharpvik" && params["id"] == "42"; ok {
		return blackbox.NewResponse().WithStatus(http.StatusOK)
	}
	return blackbox.NewResponse().WithStatus(http.StatusBadRequest)
}
