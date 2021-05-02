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
	Cookies map[string]*http.Cookie
	Body    *bytes.Buffer
}

// NewResponse returns a newly created Response instance with status 200 OK,
// and pre-allocated Headers map and Body writer.
func NewResponse() *Response {
	return &Response{
		Status:  http.StatusOK,
		Headers: make(map[string]string),
		Cookies: make(map[string]*http.Cookie),
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

// WithCookie sets cookie in Response.Cookies (overwriting any existing cookies
// with the same name) and returns reference to that same Response.
func (resp *Response) WithCookie(cookie *http.Cookie) *Response {
	resp.Cookies[cookie.Name] = cookie
	return resp
}

// Write writes given bytes to Response.Body and returns reference to the
// Response.
func (resp *Response) Write(b []byte) *Response {
	resp.Body.Write(b)
	return resp
}

// WriteString writes given string to Response.Body and returns reference to the
// Response.
func (resp *Response) WriteString(s string) *Response {
	resp.Body.WriteString(s)
	return resp
}

// EncodeJSON uses Go's default encoding/json library to encode data as JSON
// and write it into Request.Body.
func (resp *Response) EncodeJSON(data interface{}) error {
	resp.WithHeader("Content-Type", "application/json")
	return json.NewEncoder(resp.Body).Encode(data)
}

// DecodeJSON uses Go's default encoding/json library to decode data from
// Request.Body as JSON.
func (resp *Response) DecodeJSON(data interface{}) error {
	return json.NewDecoder(resp.Body).Decode(data)
}

// Respond writes information stored within this Response instance into the
// given net/http.ResponseWriter, thus, sending it to the client.
//
// Respond helps Response implement Responder interface.
func (resp *Response) Respond(w http.ResponseWriter) {
	resp.writeCookies(w)
	resp.writeHeaders(w)
	resp.writeStatus(w)
	resp.writeBody(w)
}

func (resp *Response) writeCookies(w http.ResponseWriter) {
	for _, cookie := range resp.Cookies {
		http.SetCookie(w, cookie)
	}
}

func (resp *Response) writeHeaders(w http.ResponseWriter) {
	for name, value := range resp.Headers {
		w.Header().Set(name, value)
	}
}

func (resp *Response) writeStatus(w http.ResponseWriter) {
	if resp.Status == 0 {
		resp.Status = http.StatusOK
	}
	w.WriteHeader(resp.Status)
}

func (resp *Response) writeBody(w http.ResponseWriter) {
	resp.Body.WriteTo(w)
}
