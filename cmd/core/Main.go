package main

import (
	"log"
	"net/http"

	"github.com/FetchWeb/Fetch/pkg/core"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var authService core.AuthService
	authService.Startup()
	http.HandleFunc("/signin", authService.Signin)
	http.HandleFunc("/signup", authService.Signup)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
