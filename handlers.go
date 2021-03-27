package blackbox

import "net/http"

// statusAndMessage returns a Handler that always responds with a given status
// and writes message to Response.Body.
func statusAndMessage(status int, message string) Handler {
	return HandlerFunc(func(*http.Request) (resp *Response) {
		resp = NewResponse().WithStatus(status)
		resp.WriteString(message)
		return
	})
}
