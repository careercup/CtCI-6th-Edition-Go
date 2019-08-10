package chapter4

func FindFirstCommonAncestor(parent, current *BTNode, p, q int) *BTNode {
	if current == nil {
		return nil
	}

	if current.Value == p {
		if isAncestor(current.Left, q) || isAncestor(current.Right, q) {
			return parent
		}
		return current
	}

	if current.Value == q {
		if isAncestor(current.Left, p) || isAncestor(current.Right, p) {
			return parent
		}
		return current
	}

	result := FindFirstCommonAncestor(current, current.Left, p, q)
	if result == nil {
		return FindFirstCommonAncestor(current, current.Right, p, q)
	}
	if result.Value == p && isAncestor(current.Right, q) {

		return current

	}
	if result.Value == q && isAncestor(current.Right, p) {
		return current

	}
	return result

}

func isAncestor(node *BTNode, target int) bool {
	if node == nil {
		return false
	}
	if node.Value == target {
		return true
	}
	foundInLeft := isAncestor(node.Left, target)
	if foundInLeft {
		return true
	}
	return isAncestor(node.Right, target)

}
