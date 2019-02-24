package chapter4

import (
	"reflect"
	"testing"
)

func TestMinimalTree(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out *BTNode
	}{

		"Should return nil when slice is empty": {
			in:  []int{},
			out: nil,
		},

		"Should return only root node when slice has single value": {
			in:  []int{1},
			out: &BTNode{Value: 1},
		},

		"Should return minimal BST when slice has 2 values": {
			in: []int{1, 2},
			out: &BTNode{
				Value: 1,
				Right: &BTNode{Value: 2},
			},
		},

		"Should return minimal BST when slice has 3 values": {
			in: []int{1, 2, 3},
			out: &BTNode{
				Value: 2,
				Left:  &BTNode{Value: 1},
				Right: &BTNode{Value: 3},
			},
		},

		"Should return minimal BST when slice has 4 values": {
			in: []int{1, 2, 3, 4},
			out: &BTNode{
				Value: 2,
				Left:  &BTNode{Value: 1},
				Right: &BTNode{Value: 3,
					Right: &BTNode{Value: 4},
				},
			},
		},

		"Should return minimal BST when slice has 5 values": {
			in: []int{1, 2, 3, 4, 5},
			out: &BTNode{
				Value: 3,
				Left: &BTNode{Value: 1,
					Right: &BTNode{Value: 2},
				},
				Right: &BTNode{Value: 4,
					Right: &BTNode{Value: 5},
				},
			},
		},

		"Should return minimal BST when slice has 6 values": {
			in: []int{1, 2, 3, 4, 5, 6},
			out: &BTNode{
				Value: 3,
				Left: &BTNode{Value: 1,
					Right: &BTNode{Value: 2},
				},
				Right: &BTNode{Value: 5,
					Left:  &BTNode{Value: 4},
					Right: &BTNode{Value: 6},
				},
			},
		},

		"Should return minimal BST when slice has 7 values": {
			in: []int{1, 2, 3, 4, 5, 6, 7},
			out: &BTNode{
				Value: 4,
				Left: &BTNode{Value: 2,
					Left:  &BTNode{Value: 1},
					Right: &BTNode{Value: 3},
				},
				Right: &BTNode{Value: 6,
					Left:  &BTNode{Value: 5},
					Right: &BTNode{Value: 7},
				},
			},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := MinimalTree(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%+v expected=%+v", out, test.out)
			}
		})
	}
}
