package chapter4

import (
	"reflect"
	"testing"
)

func TestSuccessor(t *testing.T) {
	// Initiate Node for parent traversal
	n1 := &BTNode{}
	n2 := &BTNode{}
	n3 := &BTNode{}
	n2.SetRight(n1)
	n3.SetLeft(n2)

	n4 := &BTNode{}
	n5 := &BTNode{}
	n6 := &BTNode{}
	n5.SetRight(n4)
	n6.SetRight(n5)

	tests := map[string]struct {
		in  *BTNode
		out *BTNode
	}{

		"Should return nil if input is nil": {
			in:  nil,
			out: nil,
		},

		"Should return nil if input doesn't have parent and children": {
			in:  &BTNode{},
			out: nil,
		},

		"Should return most left node of right subtree": {
			in: &BTNode{
				Right: &BTNode{
					Left: &BTNode{
						Value: 5,
						Left: &BTNode{
							Value: 10,
							Right: &BTNode{},
						},
					},
				},
			},
			out: &BTNode{
				Value: 10,
				Right: &BTNode{},
			},
		},

		/*
		 *     n3
		 *    /
		 *   n2
		 *    \
		 *     n1
		 */
		"Should return parent which has left child": {
			in:  n1,
			out: n3,
		},

		/*
		 *  n6
		 *    \
		 *    n5
		 *      \
		 *      n4
		 */
		"Should return nil whne input is the last node to traversal": {
			in:  n4,
			out: nil,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := Successor(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%+v expected=%+v", out, test.out)
			}
		})
	}
}
