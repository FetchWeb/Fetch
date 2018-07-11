package handlers

import (
	"fmt"
	"net/http"
)

// Debug handler
func Debug(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Go")

	s := "{\"response\": \"Hello, World! Your IP is " + r.RemoteAddr + "\"}"

	var _resp []byte

	copy(_resp[:], s)
	w.Write(_resp)
	// r.ParseForm() // parse arguments, you have to call this by yourself

	// fmt.Println(r.Form) // print form information in server side
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }
	// fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}
