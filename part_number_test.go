package magicconch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIntPtr(t *testing.T) {
	Float32Ptr(1)
	Float64Ptr(1)
	IntPtr(1)
	Int8Ptr(1)
	Int16Ptr(1)
	Int32Ptr(1)
	Int64Ptr(1)
	UintPtr(1)
	Uint8Ptr(1)
	Uint16Ptr(1)
	Uint32Ptr(1)
	Uint64Ptr(1)
}

func TestIntInSlice(t *testing.T) {
	assertions := require.New(t)
	assertions.False(IntInSlice(1, []int{}))
	assertions.False(IntInSlice(1, []int{2}))
	assertions.True(IntInSlice(1, []int{1, 2}))
}

func TestUintInSlice(t *testing.T) {
	assertions := require.New(t)
	assertions.False(UintInSlice(uint(1), []uint{}))
	assertions.False(UintInSlice(uint(1), []uint{uint(2)}))
	assertions.True(UintInSlice(uint(1), []uint{uint(1), uint(2)}))
}
