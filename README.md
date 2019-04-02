# Go Webserver

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/fetchweb/fetch)](https://goreportcard.com/report/github.com/fetchweb/fetch)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/fetchweb/fetch)
[![Build Status](https://travis-ci.org/FetchWeb/Fetch.svg?branch=master)](https://travis-ci.org/FetchWeb/Fetch)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/FetchWeb/Fetch/releases/tag/v0.1.0)

This is an experimental web server written in Go. Do not use in production.

## How to get set up

Create a new Go project and put this in your `main.go` file:

```go
package main

import (
	"fmt"
	"fetch/core"
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
```go
app := core.Server{Port: "80"}
```

If you're using SSL or static file hosting, it's important to specify the projects base directory in the server struct:
```go
baseDir, _ := os.Getwd()
app := core.Server{BaseDir: baseDir}
```
