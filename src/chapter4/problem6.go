package chapter4

// Successor checks the next node in binary tree with in-order traversal
func Successor(node *BTNode) *BTNode {
	if node == nil {
		return nil
	}

	// node has right child
	if node.Right != nil {
		for n := node.Right; ; n = n.Left {
			if n.Left == nil {
				return n
			}
		}
	}

	// node is right child of Parent
	for ; node.Parent != nil; node = node.Parent {
		if node.Parent.Left == node {
			return node.Parent
		}
	}

	return nil
}
