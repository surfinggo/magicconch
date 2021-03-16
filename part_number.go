package magicconch

// Float32Ptr is an alias to PtrToFloat32
var Float32Ptr = PtrToFloat32

// Float64Ptr is an alias to PtrToFloat64
var Float64Ptr = PtrToFloat64

// IntPtr is an alias to PtrToInt
var IntPtr = PtrToInt

// Int8Ptr is an alias to PtrToInt8
var Int8Ptr = PtrToInt8

// Int16Ptr is an alias to PtrToInt16
var Int16Ptr = PtrToInt16

// Int32Ptr is an alias to PtrToInt32
var Int32Ptr = PtrToInt32

// Int64Ptr is an alias to PtrToInt64
var Int64Ptr = PtrToInt64

// UintPtr is an alias to PtrToUint
var UintPtr = PtrToUint

// Uint8Ptr is an alias to PtrToUint8
var Uint8Ptr = PtrToUint8

// Uint16Ptr is an alias to PtrToUint16
var Uint16Ptr = PtrToUint16

// Uint32Ptr is an alias to PtrToUint32
var Uint32Ptr = PtrToUint32

// Uint64Ptr is an alias to PtrToUint64
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
