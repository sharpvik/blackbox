package middleware

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/sharpvik/blackbox"
)

// Param interface represents a named URL parameter that can be matched with a
// given regex expression. In this package, we've declared a few struct's that
// implement this interface and provide useful functionality. However, feel free
// to write your own implementations and use them to enhance your routing
// experience.
type Param interface {
	// Name returns a string that represents parameter name. Name must not be
	// empty or consist of whitespace. Your router will panic on initialisation
	// if that happens.
	Name() string

	// Regex returns a compiled regex expression. You should consider using
	// regexp.MustCompile to implement this method. MustCompile will panic on
	// invalid input and you'll be able to catch your mistake early.
	Regex() *regexp.Regexp
}

// RegexParam represents an unnamed constant regex pattern. You can use it to
// match expressions with a custom regex pattern. Though, you will not be able
// to retrieve their values from within a handler -- if you want to have access
// to the value, check out RegexParam!
type ConstParam struct {
	pattern string
}

// ParamConst accepts a regex pattern.
func ParamConst(pattern string) *ConstParam {
	return &ConstParam{pattern}
}

// Name helps ConstParam implement Param.
func (cp *ConstParam) Name() string {
	return ""
}

// Regex helps ConstParam implement Param.
func (cp *ConstParam) Regex() *regexp.Regexp {
	return regexp.MustCompile(cp.pattern)
}

// RegexParam represents a constant regex pattern. You can use it to match
// expressions with a custom regex pattern. You will then be able to access
// the matched string value through the GetParams function.
type RegexParam struct {
	name    string
	pattern string
}

// ParamConst accepts a name, and a regex pattern.
func ParamRegex(name, pattern string) *RegexParam {
	return &RegexParam{name, pattern}
}

// Name helps RegexParam implement Param.
func (cp *RegexParam) Name() string {
	return cp.name
}

// Regex helps RegexParam implement Param.
func (cp *RegexParam) Regex() *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`(?P<%s>%s)`, cp.Name(), cp.pattern))
}

// ParamString accepts a name. It represents a string parameter that can
// consist only of latin letters, digits, underscores, and whitespace.
func ParamString(name string) *RegexParam {
	return ParamRegex(name, `[\w\s]+`)
}

// ParamInt accepts a name. It represents an integer from negative infinity to
// positive infinity.
func ParamInt(name string) *RegexParam {
	return ParamRegex(name, `-?[1-9]\d*|0`)
}

// Params constructs ParamsRoute from a list of Param's and a blackbox.Handler.
func Params(params paramsList, h blackbox.Handler) *Middleware {
	return NewMiddleware(
		params.filter(),
		blackbox.HandlerFunc(func(r *http.Request) *blackbox.Response {
			return h.Handle(insertParams(params.regex(), r))
		}),
	)
}

// ParamsList is a constructor function for the internal type paramsList.
// Use it when passing Param's to the Params middleware constructor like so:
//
//     api := Params(
//             ParamsList(
//                     ParamConst("api"),
//                     ParamString("handle"),
//                     ParamInt("id")),
//             myMagicHandler)
//
func ParamsList(params ...Param) paramsList {
	return params
}

// GetParams accepts an http.Request and retrieves a map of parameters from it.
// Use this function within a handler that your pass to the Params constructor.
func GetParams(r *http.Request) map[string]string {
	return r.Context().Value(varsKey).(map[string]string)
}

func insertParams(expr *regexp.Regexp, r *http.Request) *http.Request {
	if vars := parseParams(expr, r.URL.String()); vars != nil {
		return r.WithContext(context.WithValue(r.Context(), varsKey, vars))
	}
	return r
}

func parseParams(expr *regexp.Regexp, uri string) (vars map[string]string) {
	if expr.MatchString(uri) {
		vars = varsMap(expr, expr.FindStringSubmatch(uri))
	}
	return
}

func varsMap(expr *regexp.Regexp, matches []string) (vars map[string]string) {
	vars = make(map[string]string)
	for _, name := range expr.SubexpNames() {
		if name != "" {
			vars[name] = matches[expr.SubexpIndex(name)]
		}
	}
	return
}

func (params paramsList) filter() blackbox.FilterFunc {
	expr := params.regex()
	return func(r *http.Request) bool {
		uri := r.URL.String()
		return len(expr.FindString(uri)) == len(uri)
	}
}

func (params paramsList) regex() *regexp.Regexp {
	return regexp.MustCompile(params.expr())
}

func (params paramsList) expr() string {
	var builder strings.Builder
	for _, p := range params {
		builder.WriteByte('/')
		builder.WriteString(p.Regex().String())
	}
	return builder.String()
}

type paramsList []Param

type contextKey int

const varsKey contextKey = iota
