package chapter4

// BTNode represents a node in binary tree
type BTNode struct {
	Value  int
	Parent *BTNode
	Left   *BTNode
	Right  *BTNode
}

// SetRight sets child on right side
func (n *BTNode) SetRight(c *BTNode) {
	n.Right = c
	c.Parent = n
}

// SetLeft sets child on right side
func (n *BTNode) SetLeft(c *BTNode) {
	n.Left = c
	c.Parent = n
}
