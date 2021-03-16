package magicconch

var Float32Ptr = PtrToFloat32
var Float64Ptr = PtrToFloat64
var IntPtr = PtrToInt
var Int8Ptr = PtrToInt8
var Int16Ptr = PtrToInt16
var Int32Ptr = PtrToInt32
var Int64Ptr = PtrToInt64
var UintPtr = PtrToUint
var Uint8Ptr = PtrToUint8
var Uint16Ptr = PtrToUint16
var Uint32Ptr = PtrToUint32
var Uint64Ptr = PtrToUint64

// IntInSlice determines whether an int is in a slice or not.
func IntInSlice(i int, slice []int) bool {
	for _, s := range slice {
		if s == i {
			return true
		}
	}
	return false
}

// UintInSlice determines whether an uint is in a slice or not.
func UintInSlice(i uint, slice []uint) bool {
	for _, s := range slice {
		if s == i {
			return true
		}
	}
	return false
}
