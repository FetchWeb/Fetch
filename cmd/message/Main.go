//package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/FetchWeb/Fetch/pkg/message"
// )

// func main() {
// 	// Startup the message service.
// 	var s message.Service
// 	if err := s.Startup(); err != nil {
// 		panic(err)
// 	}

// 	// Start checking for queued emails.
// 	go s.Check()

// 	// Start listening at endpoint for incoming emails.
// 	http.HandleFunc("/listen", s.Listen)
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }
