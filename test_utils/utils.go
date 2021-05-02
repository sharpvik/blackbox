package test_utils

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sharpvik/blackbox"
)

func GetAndCheckStatusAndBody(
	t *testing.T, rtr *blackbox.Router, uri string, status int, body string) (
	resp *blackbox.Response) {
	req := Get(t, uri)
	resp = rtr.Handle(req)
	assert.Equal(t, status, resp.Status, "unexpected response status")
	assert.Equal(t, body, resp.Body.String(), "unexpected response body")
	return
}

func Get(t *testing.T, uri string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	assert.NoError(t, err, "unexpected error creating request")
	return req
}

func Post(t *testing.T, uri string, body []byte) *http.Request {
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(body))
	assert.NoError(t, err, "unexpected error creating request")
	return req
}

func RespondWithStatusAndMessage(
	status int, message string) blackbox.HandlerFunc {
	return func(*http.Request) *blackbox.Response {
		resp := blackbox.NewResponse().WithStatus(status)
		resp.WriteString(message)
		return resp
	}
}

func RespondWithStatusAndCookie(
	status int, name, value string) blackbox.HandlerFunc {
	return func(*http.Request) *blackbox.Response {
		return blackbox.NewResponse().
			WithStatus(status).
			WithCookie(&http.Cookie{
				Name:  name,
				Value: value,
			})
	}
}
