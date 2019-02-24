package chapter4

import (
	"container/list"
	"reflect"
	"testing"
)

func TestListOfDepth(t *testing.T) {
	tests := map[string]struct {
		in  *BTNode
		out []*list.List
	}{

		"Should return slice of a single list when binary tree has only root": {
			in: &BTNode{
				Value: 1,
			},
			out: []*list.List{
				pushList(list.New(), &BTNode{Value: 1}),
			},
		},

		"Should return slice of lists when binary tree has nodes": {
			in: &BTNode{
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
			out: []*list.List{
				pushList(list.New(),
					&BTNode{
						Value: 4,
						Left: &BTNode{Value: 2,
							Left:  &BTNode{Value: 1},
							Right: &BTNode{Value: 3},
						},
						Right: &BTNode{Value: 6,
							Left:  &BTNode{Value: 5},
							Right: &BTNode{Value: 7},
						},
					}),
				pushList(list.New(),
					&BTNode{Value: 2,
						Left:  &BTNode{Value: 1},
						Right: &BTNode{Value: 3},
					},
					&BTNode{Value: 6,
						Left:  &BTNode{Value: 5},
						Right: &BTNode{Value: 7},
					},
				),
				pushList(list.New(),
					&BTNode{Value: 1},
					&BTNode{Value: 3},
					&BTNode{Value: 5},
					&BTNode{Value: 7},
				),
			},
		},

		"Should return slice of lists when binary tree is unbalanced": {
			in: &BTNode{
				Value: 4,
				Left:  &BTNode{Value: 2},
				Right: &BTNode{Value: 6,
					Left: &BTNode{Value: 5},
					Right: &BTNode{Value: 7,
						Right: &BTNode{Value: 8,
							Right: &BTNode{Value: 9},
						},
					},
				},
			},
			out: []*list.List{
				pushList(list.New(),
					&BTNode{
						Value: 4,
						Left:  &BTNode{Value: 2},
						Right: &BTNode{Value: 6,
							Left: &BTNode{Value: 5},
							Right: &BTNode{Value: 7,
								Right: &BTNode{Value: 8,
									Right: &BTNode{Value: 9},
								},
							},
						},
					}),
				pushList(list.New(),
					&BTNode{Value: 2},
					&BTNode{Value: 6,
						Left: &BTNode{Value: 5},
						Right: &BTNode{Value: 7,
							Right: &BTNode{Value: 8,
								Right: &BTNode{Value: 9},
							},
						},
					},
				),
				pushList(list.New(),
					&BTNode{Value: 5},
					&BTNode{Value: 7,
						Right: &BTNode{Value: 8,
							Right: &BTNode{Value: 9},
						},
					},
				),
				pushList(list.New(),
					&BTNode{Value: 8,
						Right: &BTNode{Value: 9},
					},
				),
				pushList(list.New(), &BTNode{Value: 9}),
			},
		},

		"Should return nil if root is nil": {
			in:  nil,
			out: []*list.List(nil),
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := ListOfDepth(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%+v expected=%+v", out, test.out)
			}
		})
	}
}

func pushList(l *list.List, data ...interface{}) *list.List {
	for _, d := range data {
		l.PushFront(d)
	}

	return l
}
