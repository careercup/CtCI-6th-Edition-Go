package chapter4

import "container/list"

// ListOfDepth returns slice of list which has nodes on each depth
func ListOfDepth(root *BTNode) []*list.List {
	if root == nil {
		return nil
	}

	var res []*list.List
	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		res = append(res, cloneList(queue))
		nextQueue := list.New()
		for queue.Len() > 0 {
			elm := queue.Back()
			queue.Remove(elm)
			btNode := elm.Value.(*BTNode)

			if btNode.Left != nil {
				nextQueue.PushFront(btNode.Left)
			}
			if btNode.Right != nil {
				nextQueue.PushFront(btNode.Right)
			}
		}

		queue = nextQueue
	}

	return res
}

func cloneList(l *list.List) *list.List {
	clone := list.New()
	clone.PushBackList(l)
	return clone
}
