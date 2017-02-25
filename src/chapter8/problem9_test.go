package chapter8

import (
	"reflect"
	"testing"
)

func TestParens(t *testing.T) {
	tests := map[string]struct {
		in  int
		out []string
	}{
		"Input is -1": {
			in:  -1,
			out: []string(nil),
		},
		"Input is 0": {
			in: 0,
			out: []string{
				"",
			},
		},
		"Input is 1": {
			in: 1,
			out: []string{
				"()",
			},
		},
		"Input is 2": {
			in: 2,
			out: []string{
				"(())",
				"()()",
			},
		},
		"Input is 3": {
			in: 3,
			out: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := Parens(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%+v expected=%+v", out, test.out)
			}
		})
	}
}
