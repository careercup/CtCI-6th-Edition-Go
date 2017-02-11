package chapter4

import "testing"

func TestRouteBetweenNodes(t *testing.T) {
	// Initiate Node for Graph
	node1 := &GraphNode{Value: 1, Nodes: []*GraphNode(nil)}
	node2 := &GraphNode{Value: 2, Nodes: []*GraphNode(nil)}
	node3 := &GraphNode{Value: 3, Nodes: []*GraphNode(nil)}
	node4 := &GraphNode{Value: 4, Nodes: []*GraphNode(nil)}
	node5 := &GraphNode{Value: 5, Nodes: []*GraphNode(nil)}
	node6 := &GraphNode{Value: 6, Nodes: []*GraphNode(nil)}
	node7 := &GraphNode{Value: 7, Nodes: []*GraphNode(nil)}

	// Conect nodes
	/*
	 * // un-directed
	 * 1 - 2
	 * | /
	 * 3 - 4 - 5
	 *
	 * // directed
	 * 6 -> 7
	 */
	node1.PushNodes(node2, node3)
	node2.PushNodes(node1, node3)
	node3.PushNodes(node2, node4)
	node4.PushNodes(node3, node5)
	node5.PushNodes(node4)
	node6.PushNodes(node7)
	node7.PushNodes()

	type (
		in struct {
			start *GraphNode
			end   *GraphNode
		}
	)
	tests := map[string]struct {
		in
		out bool
	}{

		"Should return true when there is a route": {
			in:  in{start: node1, end: node2},
			out: true,
		},

		"Should return true when there is a route and route is greather than 1": {
			in:  in{start: node1, end: node5},
			out: true,
		},

		"Should return true when start and end are the same": {
			in:  in{start: node1, end: node1},
			out: true,
		},

		"Should return false when there is no route": {
			in:  in{start: node1, end: node6},
			out: false,
		},

		"Should return false when end has a route to start but start deons't have a route to end": {
			in:  in{start: node7, end: node6},
			out: false,
		},

		// Edge cases
		"Should return false when start is nil": {
			in:  in{start: nil, end: node1},
			out: false,
		},

		"Should return false when end is nil": {
			in:  in{start: node1, end: nil},
			out: false,
		},

		"Should return false when start and end are nil": {
			in:  in{start: nil, end: nil},
			out: false,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := RouteBetweenNodes(test.in.start, test.in.end)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
