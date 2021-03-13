package conch

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	assertions := require.New(t)

	Must(os.Unsetenv("TEST_BLAH_BLAH"))
	assertions.Equal("no", Getenv("TEST_BLAH_BLAH", "no"))

	Must(os.Setenv("TEST_BLAH_BLAH", ""))
	assertions.Equal("no", Getenv("TEST_BLAH_BLAH", "no"))

	Must(os.Setenv("TEST_BLAH_BLAH", "yes"))
	assertions.Equal("yes", Getenv("TEST_BLAH_BLAH", "no"))
}

func TestLookupEnv(t *testing.T) {
	assertions := require.New(t)

	Must(os.Unsetenv("TEST_BLAH_BLAH"))
	assertions.Equal("no", LookupEnv("TEST_BLAH_BLAH", "no"))

	Must(os.Setenv("TEST_BLAH_BLAH", ""))
	assertions.Equal("", LookupEnv("TEST_BLAH_BLAH", "no"))

	Must(os.Setenv("TEST_BLAH_BLAH", "yes"))
	assertions.Equal("yes", LookupEnv("TEST_BLAH_BLAH", "no"))
}
