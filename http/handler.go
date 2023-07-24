package http

import (
	"context"
	"net/http"
	"prodocutApi/http/controller"
	"regexp"
	"strings"
)

type server struct {
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	newRoute("GET", `\/api\/product`, controller.GETHandlerProduct),
	newRoute("GET", `\/api\/product\/([0-9]+)`, controller.GETHandlerOneProduct),
	newRoute("POST", `\/api\/product`, controller.POSTHandlerProduct),
	newRoute("PUT", `\/api\/product`, controller.PUTHandlerProduct),
	newRoute("DELETE", `\/api\/product`, controller.DELETEHandlerProduct),
}

func newRoute(method, parttern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile(`^` + parttern + `$`), handler}
}

func NewHandler() *server {
	return &server{}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var allow []string

	for _, route := range routes {

		matches := route.regex.FindStringSubmatch(r.URL.Path)

		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), controller.CtxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
