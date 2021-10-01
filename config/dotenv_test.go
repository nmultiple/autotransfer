package config

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	testcases := []struct {
		name string
		line string
		expectError error
		expectKey string
		expectValue string
	}{
		{
			name: "normal",
			line: "foo=bar",
			expectError: nil,
			expectKey: "foo",
			expectValue: "bar",
		},
		{
			name: "multiple equal",
			line: "foo=bar=baz",
			expectError: nil,
			expectKey: "foo",
			expectValue: "bar=baz",
		},
		{
			name: "empty",
			line: "=",
			expectError: nil,
			expectKey: "",
			expectValue: "",
		},
		{
			name: "indented",
			line: "    foo=bar",
			expectError: nil,
			expectKey: "foo",
			expectValue: "bar",
		},
		{
			name: "trailing space",
			line: "foo=bar ",
			expectError: nil,
			expectKey: "foo",
			expectValue: "bar",
		},
		{
			name: "empty",
			line: "",
			expectError: EmptyLine,
		},
		{
			name: "only space",
			line: " ",
			expectError: EmptyLine,
		},
		{
			name: "no equal",
			line: "foo",
			expectError: SyntaxError,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			k, v, err := parseLine(testcase.line)
			if err != testcase.expectError {
				t.Errorf("unexpected error: expected=%v, actual=%v\n",
					testcase.expectError, err)
			}
			if k != testcase.expectKey {
				t.Errorf("unexpected key: expected=%v, actual=%v\n",
					testcase.expectKey, k)
			}
			if k != testcase.expectKey {
				t.Errorf("unexpected value: expected=%v, actual=%v\n",
					testcase.expectValue, v)
			}
		})
	}
}
