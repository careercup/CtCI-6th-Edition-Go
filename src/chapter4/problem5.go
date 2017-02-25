package chapter4

// ValidateBST checks binary tree to see if it's valid binary search tree
func ValidateBST(node *BTNode) bool {
	if node == nil {
		return true
	}
	return ValidateBSTR(node.Left, nil, node.Value) && ValidateBSTR(node.Right, node.Value, nil)
}

// ValidateBSTR checks binary tree to see if it's valid binary search tree
func ValidateBSTR(node *BTNode, min, max interface{}) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Value < min.(int) {
		return false
	}

	if max != nil && node.Value > max.(int) {
		return false
	}

	return ValidateBSTR(node.Left, min, node.Value) && ValidateBSTR(node.Right, node.Value, max)
}
