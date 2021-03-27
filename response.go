package blackbox

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Response represents a standard HTTP response.
type Response struct {
	// See https://pkg.go.dev/net/http#pkg-constants for status constants.
	Status  int
	Headers map[string]string
	Body    *bytes.Buffer
}

// NewResponse returns a newly created Response instance with status 200 OK,
// and pre-allocated Headers map and Body writer.
func NewResponse() *Response {
	return &Response{
		Status:  http.StatusOK,
		Headers: make(map[string]string),
		Body:    bytes.NewBuffer(make([]byte, 0, 1024)),
	}
}

// WithStatus sets Response status and returns reference to that same Response.
func (resp *Response) WithStatus(status int) *Response {
	resp.Status = status
	return resp
}

// WithHeader sets header in Response.Headers (overwriting any exising headers
// with the same name) and returns reference to that same Response.
func (resp *Response) WithHeader(name, value string) *Response {
	resp.Headers[name] = value
	return resp
}

// Write writes given bytes to Response.Body.
func (resp *Response) Write(b []byte) {
	resp.Body.Write(b)
}

// WriteString writes given string to Response.Body.
func (resp *Response) WriteString(s string) {
	resp.Body.WriteString(s)
}

// EncodeJSON uses Go's default encoding/json library to encode data as JSON
// and write it into Request.Body.
func (resp *Response) EncodeJSON(data interface{}) error {
	resp.WithHeader("Content-Type", "application/json")
	return json.NewEncoder(resp.Body).Encode(data)
}

// Respond writes information stored within this Response instance into the
// given net/http.ResponseWriter, thus, sending it to the client.
//
// Respond helps Response implement Responder interface.
func (resp *Response) Respond(w http.ResponseWriter) {
	resp.writeHeaders(w)
	resp.writeStatus(w)
	resp.writeBody(w)
}

func (resp *Response) writeStatus(w http.ResponseWriter) {
	if resp.Status == 0 {
		resp.Status = http.StatusOK
	}
	w.WriteHeader(resp.Status)
}

func (resp *Response) writeHeaders(w http.ResponseWriter) {
	for key, value := range resp.Headers {
		w.Header().Set(key, value)
	}
}

func (resp *Response) writeBody(w http.ResponseWriter) {
	resp.Body.WriteTo(w)
}
