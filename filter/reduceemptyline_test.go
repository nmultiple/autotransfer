package filter

import (
	"testing"
)

func TestFilterUselessEmptyLine1(t *testing.T) {
	msg := []string{
		"",
		"foo",
		"",
	}

	result := filterUselessEmptyLine(msg)
	if len(result) != 1 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "foo" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
}

func TestFilterUselessEmptyLine2(t *testing.T) {
	msg := []string{
		"",
		"",
	}

	result := filterUselessEmptyLine(msg)
	if len(result) != 0 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
}

func TestFilterUselessEmptyLine3(t *testing.T) {
	msg := []string{
		"foo",
	}

	result := filterUselessEmptyLine(msg)
	if len(result) != 1 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "foo" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
}
