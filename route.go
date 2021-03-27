package blackbox

// Route is a filtering entity that has a Handler attached to it in case it
// decides to accept given net/http.Request.
type Route interface {
	Filter
	Handler
}
