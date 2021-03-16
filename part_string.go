package magicconch

import (
	"math/rand"
	"strconv"
	"strings"
)

// StringPtr is an alias to PtrToString
var StringPtr = PtrToString

// StringEnsurePrefix returns a string, which is the original string
// if it has the prefix, or the prefix will be added.
func StringEnsurePrefix(s string, cut string) string {
	if strings.HasPrefix(s, cut) {
		return s
	}
	return cut + s
}

// StringEnsureSuffix returns a string, which is the original string
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

// StringToInt8 converts a string to an int8, errors are ignored
func StringToInt8(s string) int8 {
	i, _ := strconv.ParseInt(s, 10, 8)
	return int8(i)
}

// StringToInt16 converts a string to an int16, errors are ignored
func StringToInt16(s string) int16 {
	i, _ := strconv.ParseInt(s, 10, 16)
	return int16(i)
}

// StringToInt32 converts a string to an int32, errors are ignored
func StringToInt32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int32(i)
}

// StringToInt64 converts a string to an int64, errors are ignored
func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
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

// StringToUint8 converts a string to an uint8, errors are ignored
func StringToUint8(s string) uint8 {
	i, _ := strconv.ParseUint(s, 10, 8)
	return uint8(i)
}

// StringToUint16 converts a string to an uint16, errors are ignored
func StringToUint16(s string) uint16 {
	i, _ := strconv.ParseUint(s, 10, 16)
	return uint16(i)
}

// StringToUint32 converts a string to an uint32, errors are ignored
func StringToUint32(s string) uint32 {
	i, _ := strconv.ParseUint(s, 10, 32)
	return uint32(i)
}

// StringToUint64 converts a string to an uint64, errors are ignored
func StringToUint64(s string) uint64 {
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}

const (
	// CharsetDefault is the charset for StringRand
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
