package conch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPtrToBool(t *testing.T) {
	assertions := require.New(t)
	b := true
	assertions.Equal(&b, PtrToBool(b))
}

func TestPtrToFloat32(t *testing.T) {
	assertions := require.New(t)
	i := float32(1)
	assertions.Equal(&i, PtrToFloat32(i))
}

func TestPtrToFloat64(t *testing.T) {
	assertions := require.New(t)
	i := float64(1)
	assertions.Equal(&i, PtrToFloat64(i))
}

func TestPtrToInt(t *testing.T) {
	assertions := require.New(t)
	i := 1
	assertions.Equal(&i, PtrToInt(i))
}

func TestPtrToInt8(t *testing.T) {
	assertions := require.New(t)
	i := int8(1)
	assertions.Equal(&i, PtrToInt8(i))
}

func TestPtrToInt16(t *testing.T) {
	assertions := require.New(t)
	i := int16(1)
	assertions.Equal(&i, PtrToInt16(i))
}

func TestPtrToUint(t *testing.T) {
	assertions := require.New(t)
	i := uint(1)
	assertions.Equal(&i, PtrToUint(i))
}

func TestPtrToUint8(t *testing.T) {
	assertions := require.New(t)
	i := uint8(1)
	assertions.Equal(&i, PtrToUint8(i))
}

func TestPtrToUint16(t *testing.T) {
	assertions := require.New(t)
	i := uint16(1)
	assertions.Equal(&i, PtrToUint16(i))
}

func TestPtrToUint32(t *testing.T) {
	assertions := require.New(t)
	i := uint32(1)
	assertions.Equal(&i, PtrToUint32(i))
}

func TestPtrToUint64(t *testing.T) {
	assertions := require.New(t)
	i := uint64(1)
	assertions.Equal(&i, PtrToUint64(i))
}

func TestPtrToString(t *testing.T) {
	assertions := require.New(t)
	s := "a"
	assertions.Equal(&s, PtrToString(s))
}
