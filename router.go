package fetch

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// RequestHandler type is a placeholder for http.HandlerFunc
type RequestHandler func(w Response, r Request)

// Router struct to handle route registration
type Router struct {
	GetRegister    map[string]RequestHandler
	HeadRegister   map[string]RequestHandler
	PostRegister   map[string]RequestHandler
	PutRegister    map[string]RequestHandler
	DeleteRegister map[string]RequestHandler
	OptionRegister map[string]RequestHandler
}

// NewRouter creates a new instance of the router
func NewRouter() *Router {
	_router := &Router{}
	_router.GetRegister = make(map[string]RequestHandler)
	_router.HeadRegister = make(map[string]RequestHandler)
	_router.PostRegister = make(map[string]RequestHandler)
	_router.PutRegister = make(map[string]RequestHandler)
	_router.DeleteRegister = make(map[string]RequestHandler)
	_router.OptionRegister = make(map[string]RequestHandler)
	return _router
}

// Init will set up the registers
func Init(router *Router) {
	if router.GetRegister == nil {
		router = NewRouter()
	}
}

// Get registers a GET request route
func (router *Router) Get(route string, h RequestHandler) error {
	Init(router)

	_, ok := router.GetRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	router.GetRegister[route] = h
	return nil
}

// Head registers a HEAD request route
func (router *Router) Head(route string, h RequestHandler) error {
	Init(router)

	_, ok := router.HeadRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	router.HeadRegister[route] = h
	return nil
}

// Post registers a POST request route
func (router *Router) Post(route string, h RequestHandler) error {
	Init(router)

	_, ok := router.PostRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	router.PostRegister[route] = h
	return nil
}

// Put registers a PUT request route
func (router *Router) Put(route string, h RequestHandler) error {
	Init(router)

	_, ok := router.PutRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	router.PutRegister[route] = h
	return nil
}

// Delete registers a DELETE request route
func (router *Router) Delete(route string, h RequestHandler) error {
	Init(router)

	_, ok := router.DeleteRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	router.DeleteRegister[route] = h
	return nil
}

// Options registers a Options request route
func (router *Router) Options(route string, h RequestHandler) error {
	Init(router)

	_, ok := router.OptionRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	router.OptionRegister[route] = h
	return nil
}

// SetupRoutes takes care of initialising the routes from the register
func (router *Router) SetupRoutes(mux *http.ServeMux) {
	for key, value := range router.GetRegister {
		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("GET", value))
	}
	for key, value := range router.HeadRegister {
		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("HEAD", value))
	}
	for key, value := range router.PostRegister {
		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("POST", value))
	}
	for key, value := range router.PutRegister {
		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("PUT", value))
	}
	for key, value := range router.DeleteRegister {
		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("DELETE", value))
	}
	for key, value := range router.OptionRegister {
		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("OPTION", value))
	}
}

// makeHander sets up a generic handler function for http
func makeHandler(method string, fn RequestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check that the request method matches the one specified
		if strings.ToUpper(method) != strings.ToUpper(r.Method) {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		//
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		// Log every request to console
		// @todo: Add support to log to file (?)
		fmt.Printf("[%s] %s\t%s: %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), r.RemoteAddr, r.Method, r.URL.Path)

		var response = Response{w}
		var request = Request{R: r, DB: _db}

		fn(response, request)
	}
}
