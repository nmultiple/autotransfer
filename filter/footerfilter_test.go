package filter

import (
	"testing"
)

func TestFilterFooter1(t *testing.T) {
	msg := []string{
		"foo",
		"-- ",
		"bar",
	}

	result := filterFooter(msg)
	if len(result) != 1 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "foo" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
}
