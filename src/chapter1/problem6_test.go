package chapter1

import (
	"testing"
)

func TestBasicCompress(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"aaaabbbcdddd", "a4b3c1d4"},
		{"abcd", "abcd"}, // compressed is longer, expect input
	}
	for _, c := range cases {
		actual := BasicCompress(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %s, actual: %s\n", c.input, c.expected, actual)
		}
	}
}
func TestBasicCompressB(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"aaaabbbcdddd", "a4b3c1d4"},
		{"aaaabbbbbbbbbbbcdddd", "a4b11c1d4"},
		{"aaaabbbbbbbbbbbcd", "a4b11c1d1"},
		{"abcd", "abcd"},         // compressed is longer, expect input
		{"aabbccdd", "aabbccdd"}, // compressed is longer, expect input
	}
	for _, c := range cases {
		actual := BasicCompressB(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %s, actual: %s\n", c.input, c.expected, actual)
		}
	}
}
