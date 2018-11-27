package main

import (
	"github.com/FetchWeb/Fetch/pkg/log"
)

func main() {
	log.LogDirectory = "../../logs/"
	log.Info("LOG", "This is an info message")
	log.Debug("LOG", "This is a debug message")
	log.Warning("LOG", "This is a warning message")
	log.Error("LOG", "This is an error message")
	log.Fatal("LOG", "This is a fatal message")
}
