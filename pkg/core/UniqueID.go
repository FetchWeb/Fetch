package core

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// UniqueID generates a unique id as a string from the unix time stamp and a random 64 bit integer.
func UniqueID() string {
	return strings.Join([]string{strconv.FormatInt(time.Now().Unix(), 16), strconv.FormatInt(rand.Int63(), 16)}, "")
}
