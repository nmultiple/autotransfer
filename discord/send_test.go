package discord

import (
	"testing"
)

func TestSplitToLimit(t *testing.T) {
	testcases := []struct {
		name   string
		text   string
		limit  int
		err    error
		result []string
	}{
		{
			name:  "short",
			text:  "aaa",
			limit: 100,
			err:   nil,
			result: []string{
				"aaa",
			},
		},
		{
			name:  "ascii",
			text:  "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			limit: 150,
			err:   nil,
			result: []string{
				"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
		},
		{
			name:   "split",
			text:   "aaa",
			limit:  51,
			err:    nil,
			result: []string{"a", "a", "a"},
		},
		{
			name:   "multibyte",
			text:   "あああ",
			limit:  51,
			err:    nil,
			result: []string{"あ", "あ", "あ"},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			result, err := splitToLimit(testcase.text, testcase.limit)
			if err != testcase.err {
				t.Errorf("unexpected error: expected: %v, actual: %v\n", testcase.err, err)
			}
			if len(result) != len(testcase.result) {
				t.Errorf("unexpected length: expected: %v, actual: %v\n",
					len(testcase.result), len(result))
				return
			}
			for i, e := range result {
				if e != testcase.result[i] {
					t.Errorf("unexpected result in %dth element\n", i)
				}
			}
		})
	}
}
