package snippet

import "os"

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be the default value if the variable
// is not present or the original value is empty.
// To distinguish between an empty value and an unset value, use LookupEnv.
func Getenv(key string, defaultValue string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return defaultValue
}

// LookupEnv retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be the default value and the boolean
// will be false.
func LookupEnv(key string, defaultValue string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	return defaultValue
}
