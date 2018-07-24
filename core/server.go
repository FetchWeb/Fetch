// Package core - core web server code
package core

/**
 * A fair chunk of code snippets are from https://golang.org/doc/articles/wiki/
 * To create your own self signed keys to try out HTTPS/HTTP2 (thanks to https://gist.github.com/denji/12b3a568f092ab951456):
 *    - openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650
 *    - ln -sf server.rsa.key server.key
 *    - ln -sf server.rsa.crt server.crt
 * To then create a CSR:
 *    - openssl req -new -sha256 -key server.key -out server.csr
 *    - openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 3650
 * @author Adam Buckley <adam.buckley90@gmail.com>
 */

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	// import debug handler
)

var (
	_router *Router
	_port   string
	_mux    *http.ServeMux
)

// Server is the... server
type Server struct {
}

// TODO: Finish manifest struct
type ManifestStruct struct {
	// {
	// 	"short_name": "",
	// 	"name": "",
	// 	"icons": [
	// 	  {
	// 		"src":"",
	// 		"sizes": "",
	// 		"type": ""
	// 	  }
	// 	],
	// 	"start_url": "",
	// 	"background_color": "",
	// 	"Theme_color": "",
	// 	"display": ""
	//   }
}

// RequestHandler type is a placeholder for http.HandlerFunc
type RequestHandler func(w Response, r Request)

// Setup sets up defaults
func (server *Server) Setup() {
	_mux = http.NewServeMux()
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

	// cwd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Will apparently give perfect SSL Labs score
	// cfg := &tls.Config{
	// 	MinVersion:               tls.VersionTLS12,
	// 	CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	// 	PreferServerCipherSuites: true,
	// 	CipherSuites: []uint16{
	// 		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	// 		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	// 		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	// 		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	// 	},
	// }
	// srv := &http.Server{
	// 	Addr:         ":" + _port,
	// 	Handler:      _mux,
	// 	TLSConfig:    cfg,
	// 	TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	// }

	srv := &http.Server{
		Addr:    ":" + _port, // Normally ":443"
		Handler: _mux,        //http.FileServer(http.Dir("../" + cwd)),
	}

	http2.ConfigureServer(srv, &http2.Server{})

	server.GetRouter().Get("/manifest.json", server.manifestHandler)
	server.GetRouter().SetupRoutes(_mux)

	// TODO: Enable configuration of server keys
	log.Fatal(srv.ListenAndServeTLS("C:/Users/adamb/go/src/go-webserver/server.crt", "C:/Users/adamb/go/src/go-webserver/server.key"))

	// log.Fatal(srv.ListenAndServeTLS("C:/Users/adamb/go/src/go-webserver/server.crt", "C:/Users/adamb/go/src/go-webserver/server.key"))

	// log.Fatal(http.ListenAndServe(":"+_port, nil))
}

// SetRouter sets the current router
func (server *Server) SetRouter(router *Router) {
	_router = router
}

// GetRouter returns the router
func (server *Server) GetRouter() *Router {
	return _router
}

// TODO: Make manifest struct
func (server *Server) manifestHandler(w Response, r Request) {
	w.Write([]byte("{\"short_name\": \"Go Webserver\",\"name\": \"\",\"icons\": [{\"src\":\"\",\"sizes\": \"\",\"type\": \"\"}],\"start_url\": \"\",\"background_color\": \"\",\"Theme_color\": \"\",\"display\": \"\"}"))
}
