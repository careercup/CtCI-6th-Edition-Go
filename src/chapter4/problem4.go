package chapter4

import "math"

// CheckBalanced checks if given tree is balanced
func CheckBalanced(root *BTNode) bool {
	return CheckBalancedR(root) != math.MinInt64
}

// CheckBalancedR checks if  given tree is balanced by checking depth
// It returns math.MinInt64 if it's not balanced
func CheckBalancedR(node *BTNode) int64 {
	if node == nil {
		return -1
	}

	leftD := CheckBalancedR(node.Left)
	if leftD == math.MinInt64 {
		return math.MinInt64
	}

	rightD := CheckBalancedR(node.Right)
	if rightD == math.MinInt64 {
		return math.MinInt64
	}

	if diff := leftD - rightD; diff < -1 || diff > 1 {
		return math.MinInt64
	}
	if rightD < leftD {
		return leftD + 1
	}
	return rightD + 1
}
