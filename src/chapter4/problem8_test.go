package chapter4

import "testing"

func TestFindFirstCommonAncestor(t *testing.T) {

	root := &BTNode{
		Value: 3,
		Left: &BTNode{Value: 1,
			Left:  &BTNode{Value: 7},
			Right: &BTNode{Value: 5},
		},
		Right: &BTNode{Value: 4,
			Left:  &BTNode{Value: 9},
			Right: &BTNode{Value: 8},
		},
	}

	result := FindFirstCommonAncestor(nil, root, 7, 5)
	if result.Value != 1 {
		t.Errorf("invalid result: %d", result.Value)
	}

	result = FindFirstCommonAncestor(nil, root, 9, 5)
	if result.Value != 3 {
		t.Errorf("invalid result: %d", result.Value)
	}

}
