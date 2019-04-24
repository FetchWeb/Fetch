package core

import (
	"strings"
)

// ArrayUtil generic functions
type ArrayUtil struct{}

// Contains function for ArrayUtil to generically see if the "needle" is in
// the "haystack" array. Interface types must be the same.
func (ac *ArrayUtil) Contains(haystack []interface{}, needle interface{}) bool {
	for _, search := range haystack {
		if search == needle {
			return true
		}
	}

	return false
}

// JoinStrings will combine a variadic list of strings into a single string
func JoinStrings(strs ...string) string {
	var stringsBuilder strings.Builder
	for _, str := range strs {
		stringsBuilder.WriteString(str)
	}
	return stringsBuilder.String()
}
