// Package webserver - core web server code
package fetch

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
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fetchweb/fetch/pkg/core"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"

	"golang.org/x/net/http2"
)

var (
	_router  *Router
	_mux     *http.ServeMux
	_config  interface{}
	_db      *gorm.DB
	_headers map[string]string
)

// Server is the... server
type Server struct {
	BaseDir string
	Port    string
	Config  map[string]interface{}
}

type DatabaseStruct struct {
	Database struct {
		Database string
		Driver   string
		Username string
		Password string
		Port     uint16
	}
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

// Setup sets up defaults
func (server *Server) Setup() error {
	_headers = make(map[string]string)
	_mux = http.NewServeMux()
	server.SetRouter(NewRouter())

	// Load config
	if server.BaseDir == "" {
		server.BaseDir, _ = os.Getwd()
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "/config.yml"
	}

	dat, err := ioutil.ReadFile(server.BaseDir + configPath)
	if err != nil {
		return err
	}

	db := DatabaseStruct{}
	yaml.Unmarshal(dat, &db)

	// _config := config.NewConfig()

	// fileSource := file.NewSource(
	// 	file.WithPath(server.BaseDir + configPath),
	// )
	// _config.Load(fileSource)

	// server.Config = _config.Map()

	// extract db from config into struct
	// _config.Get("database").Scan(&database)

	if db.Database.Driver == "" || db.Database.Database == "" {
		return errors.New("Missing database connection credentials")
	}

	_db, err := gorm.Open(db.Database.Driver, core.JoinStrings(db.Database.Username, ":", db.Database.Password, "@/", db.Database.Database, "?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		return err
	}

	_db.SingularTable(true)

	// defer _db.Close()

	return nil
}

// Start starts the webserver
func (server *Server) Start() {

	// Try and get port from environment, set to 3000 as default
	if server.Port == "" {
		server.Port = os.Getenv("PORT")
	}

	if server.Port == "" {
		server.Port = "443"
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
	// 	Addr:         ":" + server.Port,
	// 	Handler:      _mux,
	// 	TLSConfig:    cfg,
	// 	TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	// }

	srv := &http.Server{
		Addr:    ":" + server.Port, // Normally ":443"
		Handler: _mux,              //http.FileServer(http.Dir("../" + cwd)),
	}

	http2.ConfigureServer(srv, &http2.Server{})

	server.GetRouter().Get("/manifest.json", server.manifestHandler)
	server.GetRouter().SetupRoutes(_mux)

	log.Fatal(srv.ListenAndServeTLS(server.BaseDir+"/server.crt", server.BaseDir+"/server.key"))

}

// SetRouter sets the current router
func (server *Server) SetRouter(router *Router) {
	_router = router
	// _router.Server = server
}

// GetRouter returns the router
func (server *Server) GetRouter() *Router {
	return _router
}

// Cleanup cleans up any connections the server might have when it's terminated
func (server *Server) Cleanup() {
	if _db != nil {
		_db.Close()
	}
}

// GetDatabase gets the active database connection
func (server *Server) GetDatabase() (*gorm.DB, error) {
	if _db == nil {
		return nil, errors.New("No active database connection found")
	}

	return _db, nil
}

// TODO: Make manifest struct
func (server *Server) manifestHandler(w Response, r Request) {
	w.Write([]byte("{\"short_name\": \"Fetch\",\"name\": \"\",\"icons\": [{\"src\":\"\",\"sizes\": \"\",\"type\": \"\"}],\"start_url\": \"\",\"background_color\": \"\",\"Theme_color\": \"\",\"display\": \"\"}"))
}

// AddHeader sets a custom header
func (server *Server) AddHeader(key string, value string) {
	_headers[key] = value
}

// RemoveHeader removes a custom set header
func (server *Server) RemoveHeader(key string) error {
	if _, ok := _headers[key]; ok {
		delete(_headers, key)
		return nil
	}

	return errors.New("Header key does not exist")
}

// GetHeaders returns all custom set headers
func (server *Server) GetHeaders() map[string]string {
	return _headers
}
