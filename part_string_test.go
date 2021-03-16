package magicconch

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestStringPtr(t *testing.T) {
	StringPtr("a")
}

func TestStringEnsurePrefix(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal("https://test", StringEnsurePrefix("https://test", "https"))
	assertions.Equal("https://test", StringEnsurePrefix("test", "https://"))
}

func TestStringEnsureSuffix(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal("a.log", StringEnsureSuffix("a.log", ".log"))
	assertions.Equal("a.log", StringEnsureSuffix("a", ".log"))
}

func TestStringInSlice(t *testing.T) {
	assertions := require.New(t)
	assertions.False(StringInSlice("a", []string{}))
	assertions.False(StringInSlice("a", []string{"b"}))
	assertions.True(StringInSlice("a", []string{"a", "b"}))
}

func TestStringToInt(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(1, StringToInt("1"))
	assertions.Equal(0, StringToInt("0"))
	assertions.Equal(-1, StringToInt("-1"))
	assertions.Equal(0, StringToInt("-0"))
	assertions.Equal(0, StringToInt("a"))
}

func TestStringToInt8(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(int8(1), StringToInt8("1"))
	assertions.Equal(int8(0), StringToInt8("0"))
	assertions.Equal(int8(-1), StringToInt8("-1"))
	assertions.Equal(int8(0), StringToInt8("-0"))
	assertions.Equal(int8(0), StringToInt8("a"))
	assertions.Equal(int8(math.MaxInt8), StringToInt8(fmt.Sprintf("%d", math.MaxInt8+1)))
}

func TestStringToInt16(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(int16(1), StringToInt16("1"))
	assertions.Equal(int16(0), StringToInt16("0"))
	assertions.Equal(int16(-1), StringToInt16("-1"))
	assertions.Equal(int16(0), StringToInt16("-0"))
	assertions.Equal(int16(0), StringToInt16("a"))
	assertions.Equal(int16(math.MaxInt16), StringToInt16(fmt.Sprintf("%d", math.MaxInt16+1)))
}

func TestStringToInt32(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(int32(1), StringToInt32("1"))
	assertions.Equal(int32(0), StringToInt32("0"))
	assertions.Equal(int32(-1), StringToInt32("-1"))
	assertions.Equal(int32(0), StringToInt32("-0"))
	assertions.Equal(int32(0), StringToInt32("a"))
	assertions.Equal(int32(math.MaxInt32), StringToInt32(fmt.Sprintf("%d", math.MaxInt32+1)))
}

func TestStringToInt64(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(int64(1), StringToInt64("1"))
	assertions.Equal(int64(0), StringToInt64("0"))
	assertions.Equal(int64(-1), StringToInt64("-1"))
	assertions.Equal(int64(0), StringToInt64("-0"))
	assertions.Equal(int64(0), StringToInt64("a"))
	// math.MaxInt64 == 9223372036854775807
	assertions.Equal(int64(math.MaxInt64), StringToInt64("9223372036854775807"))
	assertions.Equal(int64(math.MaxInt64), StringToInt64("9223372036854775808"))
}

func TestStringToUint(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(uint(1), StringToUint("1"))
	assertions.Equal(uint(0), StringToUint("0"))
	assertions.Equal(uint(0), StringToUint("-1"))
	assertions.Equal(uint(0), StringToUint("-0"))
	assertions.Equal(uint(0), StringToUint("a"))
}

func TestStringToUint8(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(uint8(1), StringToUint8("1"))
	assertions.Equal(uint8(0), StringToUint8("0"))
	assertions.Equal(uint8(0), StringToUint8("-1"))
	assertions.Equal(uint8(0), StringToUint8("-0"))
	assertions.Equal(uint8(0), StringToUint8("a"))
	assertions.Equal(uint8(math.MaxUint8), StringToUint8(fmt.Sprintf("%d", math.MaxUint8+1)))
}

func TestStringToUint16(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(uint16(1), StringToUint16("1"))
	assertions.Equal(uint16(0), StringToUint16("0"))
	assertions.Equal(uint16(0), StringToUint16("-1"))
	assertions.Equal(uint16(0), StringToUint16("-0"))
	assertions.Equal(uint16(0), StringToUint16("a"))
	assertions.Equal(uint16(math.MaxUint16), StringToUint16(fmt.Sprintf("%d", math.MaxUint16+1)))
}

func TestStringToUint32(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(uint32(1), StringToUint32("1"))
	assertions.Equal(uint32(0), StringToUint32("0"))
	assertions.Equal(uint32(0), StringToUint32("-1"))
	assertions.Equal(uint32(0), StringToUint32("-0"))
	assertions.Equal(uint32(0), StringToUint32("a"))
	assertions.Equal(uint32(math.MaxUint32), StringToUint32(fmt.Sprintf("%d", math.MaxUint32+1)))
}

func TestStringToUint64(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(uint64(1), StringToUint64("1"))
	assertions.Equal(uint64(0), StringToUint64("0"))
	assertions.Equal(uint64(0), StringToUint64("-1"))
	assertions.Equal(uint64(0), StringToUint64("-0"))
	assertions.Equal(uint64(0), StringToUint64("a"))
	// math.MaxUint64 == 18446744073709551615
	assertions.Equal(uint64(math.MaxUint64), StringToUint64("18446744073709551615"))
	assertions.Equal(uint64(math.MaxUint64), StringToUint64("18446744073709551616"))
}

func TestStringRand(t *testing.T) {
	assertions := require.New(t)
	s := StringRand(8)
	assertions.Len(s, 8)
}
