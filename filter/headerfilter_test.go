package filter

import (
	"testing"
)

func TestFilterHeader1(t *testing.T) {
	msg := []string{
		"学生団体各位",
		"",
		"ほげほげ",
	}

	result := filterHeader(msg)
	if len(result) != 2 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
	if result[1] != "ほげほげ" {
		t.Fatalf("Unexpected content: %s\n", result[1])
	}
}

func TestFilterHeader2(t *testing.T) {
	msg := []string{
		"学生団体 各位",
		"",
		"ほげほげ",
	}

	result := filterHeader(msg)
	if len(result) != 2 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
	if result[1] != "ほげほげ" {
		t.Fatalf("Unexpected content: %s\n", result[1])
	}
}

func TestFilterHeader3(t *testing.T) {
	msg := []string{
		"学生団体　各位",
		"",
		"ほげほげ",
	}

	result := filterHeader(msg)
	if len(result) != 2 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
	if result[1] != "ほげほげ" {
		t.Fatalf("Unexpected content: %s\n", result[1])
	}
}

func TestFilterHeader4(t *testing.T) {
	msg := []string{
		"学生団体各位",
		"ほげほげ",
	}

	result := filterHeader(msg)
	if len(result) != 1 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "ほげほげ" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
}

func TestFilterHeader5(t *testing.T) {
	msg := []string{
		" 学生団体各位",
		"ほげほげ",
	}

	result := filterHeader(msg)
	if len(result) != 2 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != msg[0] || result[1] != msg[1] {
		t.Fatal("Unexpected content")
	}
}

func TestFilterHeader6(t *testing.T) {
	msg := []string{
		"",
		"学生団体各位",
		"ほげほげ",
	}

	result := filterHeader(msg)
	if len(result) != 1 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "ほげほげ" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
}

func TestFilterHeader7(t *testing.T) {
	msg := []string{
		"全学学生団体各位",
		"",
		"ほげほげ",
	}

	result := filterHeader(msg)
	if len(result) != 2 {
		t.Fatalf("Unexpected length: %d\n", len(result))
	}
	if result[0] != "" {
		t.Fatalf("Unexpected content: %s\n", result[0])
	}
	if result[1] != "ほげほげ" {
		t.Fatalf("Unexpected content: %s\n", result[1])
	}
}
