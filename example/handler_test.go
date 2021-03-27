package main

import (
	"net/http"
	"testing"

	"github.com/sharpvik/blackbox"
	"github.com/stretchr/testify/assert"
)

func TestMainHandler(t *testing.T) {
	getAndCheckStatusAndBody(t, mainHandler(users), "/",
		http.StatusNotImplemented, "Not Implemented")
	resp := getAndCheckStatusAndBody(t, mainHandler(users), "/api/Viktor",
		http.StatusOK, "{\"Name\":\"Viktor\",\"Age\":21}\n")
	assert.Equal(t, "application/json", resp.Headers["Content-Type"])
}

func getAndCheckStatusAndBody(
	t *testing.T, rtr *blackbox.Router, uri string, status int, body string) (
	resp *blackbox.Response) {
	req := get(t, uri)
	resp = rtr.Handle(req)
	assert.Equal(t, status, resp.Status, "unexpected response status")
	assert.Equal(t, body, resp.Body.String(), "unexpected response body")
	return
}

func get(t *testing.T, uri string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	assert.NoError(t, err, "unexpected error creating request")
	return req
}
