package main

import (
	"net/http"
	"testing"

	"github.com/sharpvik/blackbox/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestMainHandler(t *testing.T) {
	test_utils.GetAndCheckStatusAndBody(t, mainHandler(users), "/",
		http.StatusNotImplemented, "Not Implemented")
	resp := test_utils.GetAndCheckStatusAndBody(t, mainHandler(users), "/api/Viktor",
		http.StatusOK, "{\"Name\":\"Viktor\",\"Age\":21}\n")
	assert.Equal(t, "application/json", resp.Headers["Content-Type"])
}
