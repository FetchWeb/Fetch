package tests

import (
	"fetch"
	"os"
	"testing"
)

func TestServerSetup(t *testing.T) {
	app := webserver.Server{}

	app.Setup()

	currentDir, _ := os.Getwd()
	if currentDir != app.BaseDir {
		t.Error("BaseDir default not set correctly")
	}

}
