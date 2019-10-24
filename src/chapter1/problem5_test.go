package chapter1

import (
	"testing"
)

func TestAreOneEditAway(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"abcd", "abcd", true},
		{"abcd", "abcc", true},
		{"abcd", "accc", false},
		{"abcd", "abcde", true},
		{"abcd", "abcdef", false},
		{"abcde", "abcd", true},
		{"abcdef", "abcd", false},
		{"abbc", "abcd", false},
		{"abbcd", "abcd", true},
		{" ", "", true},
		{"", " ", true},
		{"", "", true},
	}
	for _, c := range cases {
		actual := AreOneEditAway(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Inputs %s, %s. Expected: %t, actual: %t\n",
				c.input1, c.input2, c.expected, actual)
		}
	}
}

func TestAreOneEditAwayB(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"abcd", "abcd", true},
		{"abcd", "abcc", true},
		{"abcd", "accc", false},
		{"abcd", "abcde", true},
		{"abcd", "abcdef", false},
		{"abcde", "abcd", true},
		{"abcdef", "abcd", false},
		{"abbc", "abcd", false},
		{"abbcd", "abcd", true},
		{" ", "", true},
		{"", " ", true},
		{"", "", true},
	}
	for _, c := range cases {
		actual := AreOneEditAwayB(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Inputs %s, %s. Expected: %t, actual: %t\n",
				c.input1, c.input2, c.expected, actual)
		}
	}
}
