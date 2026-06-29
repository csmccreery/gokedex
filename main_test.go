package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    Hello world     ",
			expected: []string{"Hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Error: %s does not match %s", word, expectedWord)
			}
		}
	}
}
