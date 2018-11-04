package core

import (
	"fmt"
	"math/rand"
	"time"
)

// UniqueID generates a unique id as a string from the unix time stamp and a random 64 bit integer.
func UniqueID() string {
	return fmt.Sprintf("%v%v", time.Now().Unix(), rand.Int63())
}
