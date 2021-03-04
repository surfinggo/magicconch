package snippet

import "testing"

func TestStringInSlice(t *testing.T) {
	if StringInSlice("a", []string{}) {
		t.Fatal("wrong answer")
	}

	if StringInSlice("a", []string{"b"}) {
		t.Fatal("wrong answer")
	}

	if !StringInSlice("a", []string{"a", "b"}) {
		t.Fatal("wrong answer")
	}
}

