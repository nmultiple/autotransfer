package filter

import (
	"testing"
)

func TestFilterQuote1(t *testing.T) {
	msg := []string{
		"foo",
		"bar",
		"On 2021/06/07 16:30, Hogetaro wrote:",
		"> foo",
		"> bar",
		"",
		"hoge",
	}

	result := filterQuote(msg)
	if len(result) != 4 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "foo" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
	if result[1] != "bar" {
		t.Fatalf("Unexpected content: %s\n", result[1])
	}
	if result[2] != "" {
		t.Fatalf("Unexpected content: %s\n", result[2])
	}
	if result[3] != "hoge" {
		t.Fatalf("Unexpected content: %s\n", result[3])
	}
}

func TestFilterQuote2(t *testing.T) {
	msg := []string{
		"On 2021/06/07 16:30, Hogetaro wrote:",
		"> foo",
		"> bar",
	}

	result := filterQuote(msg)
	if len(result) != 0 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
}
