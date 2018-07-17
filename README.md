# Go Webserver

This is an experimental web server written in Go. Do not use in production.

## How to get set up

Create a new Go project and put this in your `main.go` file:

```go
package main

import (
	"fmt"
	"go-webserver/core"
)

func main() {
	app := core.Server{}

	app.Setup()

	app.GetRouter().Get("/", homeHandler)

	app.Start()
}

func homeHandler(w core.Response, r core.Request) {
	fmt.Fprint(w, "Hello Go!\n")
}
```

You can change the port by setting the "_port" parameter on the Server struct:
```
app := core.Server{_port: "80"}
```