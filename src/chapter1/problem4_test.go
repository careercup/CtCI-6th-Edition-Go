package chapter1

import (
	"testing"
)

func TestIsPalindromePerm(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"amanaplanacanalpanama", true},   // normal palindrome
		{"amanalpancaaanplanama", true},   // jumbled palindrome
		{"amanaplanacanalpanamab", false}, // not a palindrome
		{"a", true},
		{"", true},
	}
	for _, c := range cases {
		actual := IsPalindromePerm(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %t, actual: %t\n", c.input, c.expected, actual)
		}
	}
}

func TestIsPalindromePermB(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"My rats live on no evil star", false}, // not a  palindrome
		{"Rats live on no evil star", true},     // normal palindrome
		{"amanaplanacanalpanama", true},         // normal palindrome
		{"amanalpancaaanplanama", true},         // jumbled palindrome
		{"amanaplanacanalpanamab", false},       // not a palindrome
		{"a", true},
		{"", true},
	}
	for _, c := range cases {
		actual := IsPalindromePermB(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %t, actual: %t\n", c.input, c.expected, actual)
		}
	}
}
