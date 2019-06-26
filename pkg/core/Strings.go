package core

import "strings"

func JoinStrings(strs ...string) string {
	var stringsBuilder strings.Builder
	for _, str := range strs {
		stringsBuilder.WriteString(str)
	}
	return stringsBuilder.String()
}
