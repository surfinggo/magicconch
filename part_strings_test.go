package conch

import (
	"github.com/stretchr/testify/require"

	"testing"
)

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

func TestStringToUint(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(uint(1), StringToUint("1"))
	assertions.Equal(uint(0), StringToUint("0"))
	assertions.Equal(uint(0), StringToUint("-1"))
	assertions.Equal(uint(0), StringToUint("-0"))
	assertions.Equal(uint(0), StringToUint("a"))
}

func TestStringRand(t *testing.T) {
	assertions := require.New(t)
	s := StringRand(8)
	assertions.Len(s, 8)
}
