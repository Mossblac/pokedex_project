package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},

		{
			input:    " jello by bugs  ",
			expected: []string{"jello", "by", "bugs"},
		},

		{
			input:    "Amsterdam is JUICY",
			expected: []string{"amsterdam", "is", "juicy"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual different than expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words do not match expected")
			}
		}
	}
}
