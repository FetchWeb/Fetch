// A fair chunk of code snippets are from https://golang.org/doc/articles/wiki/
package main

import (
	"log"
	"net/http"
	"os"

	// "runtime/debug"

	// import debug handler
	"go-webserver/built-in/handlers"
	"go-webserver/core"
)

// Page struct for generic pages
type Page struct {
	Title string
	Body  []byte
}

func main() {
	router := *core.NewRouter()

	// Instead of
	// http.HandleFunc("/", makeHandler(handlers.Debug))

	// We can write
	router.RegisterRoute("/", handlers.Debug)

	// Try and get port from environment, set to 3000 as default
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	router.SetupRoutes()
	log.Fatal(http.ListenAndServe(port, nil))
}
