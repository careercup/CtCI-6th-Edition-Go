package chapter4

// GraphNode represents a node in graph
type GraphNode struct {
	Value int
	Nodes []*GraphNode
}

// PushNodes append node into node list
func (n *GraphNode) PushNodes(nodes ...*GraphNode) {
	n.Nodes = append(n.Nodes, nodes...)
}
