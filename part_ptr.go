package conch

// PtrToBool creates a pointer to bool
func PtrToBool(b bool) *bool {
	return &b
}

// PtrToInt creates a pointer to int
func PtrToInt(i int) *int {
	return &i
}

// PtrToInt32 creates a pointer to int32
func PtrToInt32(i int32) *int32 {
	return &i
}

// PtrToInt64 creates a pointer to int64
func PtrToInt64(i int64) *int64 {
	return &i
}

// PtrToString creates a pointer to string
func PtrToString(s string) *string {
	return &s
}
