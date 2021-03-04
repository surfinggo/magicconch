package snippet

import (
	"strconv"
	"strings"
)

// StringEnsurePrefix returns a string, which is the original string
// if it has the prefix, or the prefix will be added.
func StringEnsurePrefix(s string, cut string) string {
	if strings.HasPrefix(s, cut) {
		return s
	}
	return cut + s
}

// StringEnsurePrefix returns a string, which is the original string
// if it has the suffix, or the suffix will be appended.
func StringEnsureSuffix(s string, cut string) string {
	if strings.HasSuffix(s, cut) {
		return s
	}
	return s + cut
}

// StringInSlice determines whether a string is in a slice or not.
func StringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// StringToInt converts a string to an int, errors are ignored
func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
