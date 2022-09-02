package customrouter

import (
	"log"
	"net/http"
	"regexp"
)

//RouteEntry is a struct that holds the method, path and handlerFunc
type RouteEntry struct {
	Path        *regexp.Regexp
	Method      string
	HandlerFunc http.HandlerFunc
}

//MyResponseWriter is a struct that holds the status code and the response writer
type MyResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

//Match is a function that returns the params from the path
func (ent *RouteEntry) Match(r *http.Request) map[string]string {
	match := ent.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		return nil // No match found
	}
	if ent.Method != r.Method {
		return nil // Method not allowed
	}
	// Create a map to store URL parameters in
	params := make(map[string]string)
	groupNames := ent.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}
	return params
}

//Router will send all requests to the router to be handled
type Router struct {
	routes []RouteEntry
}

//Route is a function that returns a router
func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	e := RouteEntry{
		Method:      method,
		Path:        regexp.MustCompile(path),
		HandlerFunc: handlerFunc,
	}
	rtr.routes = append(rtr.routes, e)
}

//ServerHTTP is a function that returns the handler for the router
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:", r) // Log the error
			http.Error(w, "{\"message\": Internal Server Error }", http.StatusInternalServerError)
		}
	}()
	for _, e := range rtr.routes {
		params := e.Match(r)
		if params == nil {
			continue // No match found
		}

		return
	}
	http.NotFound(w, r)
}
