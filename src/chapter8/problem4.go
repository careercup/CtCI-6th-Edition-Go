package chapter8

// PowerSet returns all subset of give set
func PowerSet(set []int) [][]int {
	if len(set) == 0 {
		return [][]int{
			[]int{},
		}
	}

	var res [][]int
	for _, subset := range PowerSet(set[1:]) {
		res = append(res, subset, append(subset, set[0]))
	}

	return res
}
