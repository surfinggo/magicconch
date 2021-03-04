package snippet

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPtrToStringBool(t *testing.T) {
	assertions := require.New(t)
	b := true
	assertions.Equal(&b, PtrToBool(b))
}

func TestPtrToStringInt(t *testing.T) {
	assertions := require.New(t)
	i := 1
	assertions.Equal(&i, PtrToInt(i))
}

func TestPtrToStringInt32(t *testing.T) {
	assertions := require.New(t)
	i := int32(1)
	assertions.Equal(&i, PtrToInt32(i))
}

func TestPtrToStringInt64(t *testing.T) {
	assertions := require.New(t)
	i := int64(1)
	assertions.Equal(&i, PtrToInt64(i))
}

func TestPtrToString(t *testing.T) {
	assertions := require.New(t)
	s := "a"
	assertions.Equal(&s, PtrToString(s))
}
