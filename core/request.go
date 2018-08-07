package core

import "net/http"

// Request type is a placeholder for *http.Request
type Request struct {
	R *http.Request

	Server *Server
}
