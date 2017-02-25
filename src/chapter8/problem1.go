package chapter8

// TripleStep calles TripleStepR to calculates possible ways to claim
func TripleStep(num int) int {
	if num <= 0 {
		return 0
	}
	memo := make([]int, num+1)

	return TripleStepR(num, memo)
}

// TripleStepR calculates possible ways to claim give number of stairs with 3, 2, 1 steps
func TripleStepR(num int, memo []int) int {
	if num < 0 {
		return 0
	} else if num == 0 {
		return 1
	}

	if memo[num] == 0 {
		memo[num] = TripleStepR(num-3, memo) +
			TripleStepR(num-2, memo) +
			TripleStepR(num-1, memo)
	}

	return memo[num]
}
