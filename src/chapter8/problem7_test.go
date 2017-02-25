package chapter8

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := map[string]struct {
		in  string
		out []string
	}{
		"should return slice of empty string when empty string is given": {
			in:  "",
			out: []string{""},
		},
		"should return 1 permutations when a is given": {
			in:  "a",
			out: []string{"a"},
		},
		"should return 2 permutations when ab is given": {
			in:  "ab",
			out: []string{"ab", "ba"},
		},
		"should return 6 permutations when abc is given": {
			in:  "abc",
			out: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		"should return 24 permutations when abcd is given": {
			in: "abcd",
			out: []string{
				"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
				"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
				"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
				"dabc", "dacb", "dbac", "dbca", "dcab", "dcba",
			},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := Permutations(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%+v expected=%+v", out, test.out)
			}
		})
	}
}
