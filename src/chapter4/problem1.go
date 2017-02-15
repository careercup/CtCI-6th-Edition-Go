package chapter4

import "container/list"

// RouteBetweenNodes returns true if there is a path between start and end on directed graph
func RouteBetweenNodes(start, end *GraphNode) bool {
	if start == nil || end == nil {
		return false
	}
	visits := make(map[*GraphNode]bool, 0)

	list := list.New()
	list.PushFront(start)
	for list.Len() > 0 {
		node := list.Back()
		list.Remove(node)
		gNode := node.Value.(*GraphNode)

		if gNode == end {
			return true
		}

		visits[gNode] = true
		for _, n := range gNode.Nodes {
			if _, ok := visits[n]; !ok {
				list.PushFront(n)
			}
		}
	}

	return false
}
