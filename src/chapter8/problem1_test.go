package chapter8

import "testing"

func TestTripleStep(t *testing.T) {
	tests := map[string]struct {
		in  int
		out int
	}{
		"Num is 0": {
			in:  0,
			out: 0,
		},
		"Num is 1": {
			in:  1,
			out: 1,
		},
		"Num is 2": {
			in:  2,
			out: 2,
		},
		"Num is 3": {
			in:  3,
			out: 4,
		},
		"Num is 5": {
			in:  5,
			out: 13,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := TripleStep(test.in)
			if out != test.out {
				t.Errorf("actual=%d expected=%d", out, test.out)
			}
		})
	}
}
