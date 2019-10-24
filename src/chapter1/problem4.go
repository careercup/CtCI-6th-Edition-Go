package chapter1

func IsPalindromePerm(input string) bool {
	counts := make(map[rune]int)
	inputLen := 0
	for _, r := range input {
		counts[r]++
		inputLen++
	}
	odd := inputLen%2 == 1
	seenOdd := false
	for _, val := range counts {
		if val%2 == 1 {
			if !seenOdd && odd {
				seenOdd = true
			} else {
				return false
			}
		}
	}
	return true
}

// with space O(1)
func IsPalindromePermB(input string) bool {
	bitVector := 0

	for _, r := range input {
		index := getIndexValue(r)
		if index == -1 {
			continue
		}
		mask := 1 << uint8(index)
		if (bitVector & mask) == 0 {
			bitVector |= mask
		} else {
			bitVector &= ^mask
		}
	}
	return (bitVector & (bitVector - 1)) == 0

}

func getIndexValue(r rune) int8 {

	if r < 'A' || r > 'z' {
		return int8(-1)
	}

	if r >= 'a' {
		return int8(r - 'a')
	} else if r >= 'A' {
		return int8(r - 'A')
	}

	return int8(-1)
}
