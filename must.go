package conch

// Must checks a value to be nil, panics if not
func Must(e interface{}) {
	if e != nil {
		panic(e)
	}
}
