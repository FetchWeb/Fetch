// Package core - core web server code
package core

/**
 * A fair chunk of code snippets are from https://golang.org/doc/articles/wiki/
 *
 * @author Adam Buckley <adam.buckley90@gmail.com>
 */

import (
	"log"
	"net/http"
	"os"
	// import debug handler
)

var (
	_router *Router
	_port   string
)

// Server is the... server
type Server struct {
}

// Response type is a placeholder for http.ResponseWriter
type Response http.ResponseWriter

// Request type is a placeholder for *http.Request
type Request *http.Request

// RequestHandler type is a placeholder for http.HandlerFunc
type RequestHandler func(w Response, r Request)

// Setup sets up defaults
func (server *Server) Setup() {
	server.SetRouter(NewRouter())
}

// Start starts the webserver
func (server *Server) Start() {
	// Try and get port from environment, set to 3000 as default
	if _port == "" {
		_port = os.Getenv("PORT")
	}

	if _port == "" {
		_port = "3000"
	}

	server.GetRouter().SetupRoutes()
	log.Fatal(http.ListenAndServe(":"+_port, nil))
}

// SetRouter sets the current router
func (server *Server) SetRouter(router *Router) {
	_router = router
}

// GetRouter returns the router
func (server *Server) GetRouter() *Router {
	return _router
}
