package chapter4

import "testing"

func TestCheckBalanced(t *testing.T) {
	tests := map[string]struct {
		in  *BTNode
		out bool
	}{

		"Should return true when tree only has a single node": {
			in:  &BTNode{},
			out: true,
		},

		"Should return true when difference of height of subtree is less than or equal to 1": {
			in: &BTNode{
				Right: &BTNode{
					Right: &BTNode{},
				},
				Left: &BTNode{},
			},
			out: true,
		},

		"Should return false when difference of height of subtree is greater than 1": {
			in: &BTNode{
				Right: &BTNode{
					Right: &BTNode{
						Right: &BTNode{},
						Left:  &BTNode{},
					},
				},
				Left: &BTNode{},
			},
			out: false,
		},

		"Should return true when tree is nil": {
			in:  nil,
			out: true,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := CheckBalanced(test.in)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
