package chapter4

import "testing"

func TestValidateBST(t *testing.T) {
	tests := map[string]struct {
		in  *BTNode
		out bool
	}{

		"Should return true when root is nul": {
			in:  nil,
			out: true,
		},

		"Should return true when BST has only root": {
			in:  &BTNode{},
			out: true,
		},

		"Should return true when left and right children are valid": {
			in: &BTNode{
				Value: 10,
				Left:  &BTNode{Value: 5},
				Right: &BTNode{Value: 15},
			},
			out: true,
		},

		"Should return true when left and right subtrees of root are valid binary search tree": {
			in: &BTNode{
				Value: 10,
				Left: &BTNode{
					Value: 5,
					Left:  &BTNode{Value: 1},
					Right: &BTNode{Value: 7},
				},
				Right: &BTNode{
					Value: 15,
					Left:  &BTNode{Value: 12},
					Right: &BTNode{Value: 17},
				},
			},
			out: true,
		},

		"Should return false when left child of root is invalid": {
			in: &BTNode{
				Value: 10,
				Left:  &BTNode{Value: 15},
				Right: &BTNode{Value: 15},
			},
			out: false,
		},

		"Should return false when right child of root is invalid": {
			in: &BTNode{
				Value: 10,
				Left:  &BTNode{Value: 5},
				Right: &BTNode{Value: 5},
			},
			out: false,
		},

		"Should return false when left subtrees of root is valid binary search tree but max value of the subtree is greater than root": {
			in: &BTNode{
				Value: 10,
				Left: &BTNode{
					Value: 5,
					Left:  &BTNode{Value: 1},
					Right: &BTNode{Value: 11},
				},
				Right: &BTNode{
					Value: 15,
					Left:  &BTNode{Value: 12},
					Right: &BTNode{Value: 17},
				},
			},
			out: false,
		},

		"Should return false when right subtrees of root is valid binary search tree but max value of the subtree is greater than root": {
			in: &BTNode{
				Value: 10,
				Left: &BTNode{
					Value: 5,
					Left:  &BTNode{Value: 1},
					Right: &BTNode{Value: 7},
				},
				Right: &BTNode{
					Value: 15,
					Left:  &BTNode{Value: 7},
					Right: &BTNode{Value: 17},
				},
			},
			out: false,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := ValidateBST(test.in)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
