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

type MiddlewareAdapter func(RequestHandler) http.HandlerFunc

// Router struct to handle route registration
type Router struct {
	Middleware     *RouterMiddleware
	GetRegister    map[string]func(w Response, r Request)
	HeadRegister   map[string]func(w Response, r Request)
	PostRegister   map[string]func(w Response, r Request)
	PutRegister    map[string]func(w Response, r Request)
	DeleteRegister map[string]func(w Response, r Request)
	OptionRegister map[string]func(w Response, r Request)
}

type RouterMiddleware struct {
	Executables    map[string]RequestHandler
	GetRegister    map[string][]string
	HeadRegister   map[string][]string
	PostRegister   map[string][]string
	PutRegister    map[string][]string
	DeleteRegister map[string][]string
	OptionRegister map[string][]string
}

// NewRouter creates a new instance of the router
func NewRouter() *Router {
	_router := &Router{}
	_router.GetRegister = make(map[string]func(w Response, r Request))
	_router.HeadRegister = make(map[string]func(w Response, r Request))
	_router.PostRegister = make(map[string]func(w Response, r Request))
	_router.PutRegister = make(map[string]func(w Response, r Request))
	_router.DeleteRegister = make(map[string]func(w Response, r Request))
	_router.OptionRegister = make(map[string]func(w Response, r Request))

	_router.Middleware = NewRouterMiddleware()
	return _router
}

// NewRouterMiddleware creates a new instance of the routers middleware register
func NewRouterMiddleware() *RouterMiddleware {
	_middleware := &RouterMiddleware{}
	_middleware.GetRegister = make(map[string][]string)
	_middleware.HeadRegister = make(map[string][]string)
	_middleware.PostRegister = make(map[string][]string)
	_middleware.PutRegister = make(map[string][]string)
	_middleware.DeleteRegister = make(map[string][]string)
	_middleware.OptionRegister = make(map[string][]string)

	return _middleware
}

// Init will set up the registers
func Init(router *Router) {
	if router.GetRegister == nil {
		router = NewRouter()
	}
}

// Get registers a GET request route
func (router *Router) Get(route string, m []string, h func(w Response, r Request)) error {
	Init(router)

	_, ok := router.GetRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	if len(m) > 0 {
		router.Middleware.GetRegister[route] = m
	}

	router.GetRegister[route] = h
	return nil
}

// Head registers a HEAD request route
func (router *Router) Head(route string, m []string, h func(w Response, r Request)) error {
	Init(router)

	_, ok := router.HeadRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	if len(m) > 0 {
		router.Middleware.HeadRegister[route] = m
	}

	router.HeadRegister[route] = h
	return nil
}

// Post registers a POST request route
func (router *Router) Post(route string, m []string, h func(w Response, r Request)) error {
	Init(router)

	_, ok := router.PostRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	if len(m) > 0 {
		router.Middleware.PostRegister[route] = m
	}

	router.PostRegister[route] = h
	return nil
}

// Put registers a PUT request route
func (router *Router) Put(route string, m []string, h func(w Response, r Request)) error {
	Init(router)

	_, ok := router.PutRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	if len(m) > 0 {
		router.Middleware.PutRegister[route] = m
	}

	router.PutRegister[route] = h
	return nil
}

// Delete registers a DELETE request route
func (router *Router) Delete(route string, m []string, h func(w Response, r Request)) error {
	Init(router)

	_, ok := router.DeleteRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	if len(m) > 0 {
		router.Middleware.DeleteRegister[route] = m
	}

	router.DeleteRegister[route] = h
	return nil
}

// Options registers a Options request route
func (router *Router) Options(route string, m []string, h func(w Response, r Request)) error {
	Init(router)

	_, ok := router.OptionRegister[route]
	if ok {
		return errors.New("Route \"" + route + "\" already registered")
	}

	if len(m) > 0 {
		router.Middleware.OptionRegister[route] = m
	}

	router.OptionRegister[route] = h
	return nil
}

// SetupRoutes takes care of initialising the routes from the register
func (router *Router) SetupRoutes(mux *http.ServeMux) {
	for key, value := range router.GetRegister {
		fmt.Println("Setting up route: " + key)
		var m []string
		if _middleware.GetRegister[key] != nil {
			m = _middleware.GetRegister[key]
		}

		mux.HandleFunc(key, makeHandler("GET", value, m))
	}
	for key, value := range router.HeadRegister {
		var m []string
		if _middleware.HeadRegister[key] != nil {
			m = _middleware.HeadRegister[key]
		}

		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("HEAD", value, m))
	}
	for key, value := range router.PostRegister {
		var m []string
		if _middleware.PostRegister[key] != nil {
			m = _middleware.PostRegister[key]
		}

		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("POST", value, m))
	}
	for key, value := range router.PutRegister {
		var m []string
		if _middleware.PutRegister[key] != nil {
			m = _middleware.PutRegister[key]
		}

		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("PUT", value, m))
	}
	for key, value := range router.DeleteRegister {
		var m []string
		if _middleware.DeleteRegister[key] != nil {
			m = _middleware.DeleteRegister[key]
		}

		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("DELETE", value, m))
	}
	for key, value := range router.OptionRegister {
		var m []string
		if _middleware.OptionRegister[key] != nil {
			m = _middleware.OptionRegister[key]
		}

		fmt.Println("Setting up route: " + key)
		mux.HandleFunc(key, makeHandler("OPTION", value, m))
	}
}

// makeHander sets up a generic handler function for http
func makeHandler(method string, fn func(w Response, r Request), m []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check that the request method matches the one specified
		if strings.ToUpper(method) != strings.ToUpper(r.Method) {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		//
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		if len(_headers) > 0 {
			for key, value := range _headers {
				w.Header().Set(key, value)
			}
		}
		// Log every request to console
		// @todo: Add support to log to file (?)
		fmt.Printf("[%s] %s\t%s: %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), r.RemoteAddr, r.Method, r.URL.Path)

		var response = Response{w}
		var request = Request{R: r, DB: _db}

		// if len(m) > 0 {

		// 	var returnFunction RequestHandler

		// 	returnFunction = MiddlewareAdapter(fn(response, request))
		// 	// Evaluate middleware layers if given
		// 	for middlewareExecIndex, middlewareFuncString := range m {
		// 		fmt.Printf("Middleware looking for %s in position %d", middlewareFuncString, middlewareExecIndex)

		// 		if _middleware.Executables[middlewareFuncString] != nil {
		// 			fmt.Println("Found middleware function")
		// 			returnFunction = MiddlewareAdapter(_middleware.Executables[middlewareFuncString])
		// 		}
		// 	}

		// 	returnFunction = MiddlewareAdapter(returnFunction)
		// } else {
		fn(response, request)
		// }
	}
}
