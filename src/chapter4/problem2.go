package chapter4

// MinimalTree creates a binary search tree by given sorted slice
func MinimalTree(nums []int) *BTNode {
	if len(nums) == 0 {
		return nil
	}

	return MinimalTreeR(nums, 0, len(nums)-1)
}

// MinimalTreeR is recursive function to build minimal tree
func MinimalTreeR(nums []int, start, end int) *BTNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	left := MinimalTreeR(nums, start, mid-1)
	right := MinimalTreeR(nums, mid+1, end)

	return &BTNode{
		Value: nums[mid],
		Left:  left,
		Right: right,
	}
}
