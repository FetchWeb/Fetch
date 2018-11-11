package main

import (
	"log"
	"net/http"

	"github.com/FetchWeb/Fetch/pkg/message"
)

func main() {
	var s message.Service
	if err := s.Startup(); err != nil {
		panic(err)
	}

	go s.Check()
	http.HandleFunc("/listen", s.Listen)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
