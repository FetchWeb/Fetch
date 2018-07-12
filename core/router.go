package core

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Router struct to handle route registration
type Router struct {
	Register map[string]func(http.ResponseWriter, *http.Request)
}

// NewRouter creates a new instance of the router
func NewRouter() *Router {
	_router := &Router{}
	_router.Register = make(map[string]func(http.ResponseWriter, *http.Request))
	return _router
}

// RegisterRoute registers a route string to a http.HandlerFunc
func (router *Router) RegisterRoute(route string, h http.HandlerFunc) error {
	if router.Register == nil {
		router = NewRouter()
	}

	_, ok := router.Register[route]
	if ok {
		return errors.New("Route already registered")
	}

	router.Register[route] = h
	return nil
}

// RetrieveRoute returns a route based on string given
func (router *Router) RetrieveRoute(route string) (http.HandlerFunc, error) {
	fn, ok := router.Register[route]
	if !ok {
		return nil, errors.New("Route \"" + route + "\" not registered")
	}

	return fn, nil
}

// SetupRoutes takes care of initialising the routes from the register
func (router *Router) SetupRoutes() {
	for key, value := range router.Register {
		fmt.Println("Setting up route: " + key)
		http.HandleFunc(key, makeHandler(value))
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log every request to console
		// @todo: Add support to log to file (?)
		fmt.Printf("[%s] %s\t%s: %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), r.RemoteAddr, r.Method, r.URL.Path)

		fn(w, r)
	}
}
