package magicconch

import (
	"math/rand"
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

// StringToUint converts a string to an uint, errors are ignored
func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	if i < 0 {
		return 0
	}
	return uint(i)
}

const (
	CharsetDefault = "abcdefghijklmnopqrstuvwxyzsABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// StringRandWithCharset generates random string with specific length and charset.
func StringRandWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// StringRand generates random string with specific length.
func StringRand(length int) string {
	return StringRandWithCharset(length, CharsetDefault)
}
