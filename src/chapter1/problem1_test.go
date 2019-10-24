package chapter1

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abcc", false},
		{"asbckmdsf", false},
		{"  ", false},
		{" ", true},
		{"", true},
	}
	for _, c := range cases {
		actual := IsUnique(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %t, actual: %t\n", c.input, c.expected, actual)
		}
	}
}

func TestIsUniqueB(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abcc", false},
		{"asbckmdsf", false},
	}
	for _, c := range cases {
		actual := IsUniqueB(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %t, actual: %t\n", c.input, c.expected, actual)
		}
	}
}
