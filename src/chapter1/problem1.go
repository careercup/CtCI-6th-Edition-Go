package chapter1

// Using map of runes for duplicate detection.
func IsUnique(input string) bool {
	seen := make(map[rune]struct{})
	for _, r := range input {
		if _, ok := seen[r]; ok {
			return false
		} else {
			seen[r] = struct{}{}
		}
	}
	return true
}

// Using no additional data structures. Assumes the string only contains alphabet.
func IsUniqueB(input string) bool {
	seen := 0
	for _, r := range input {
		rBit := 1 << uint32(r-'A')
		if (seen & rBit) != 0 {
			return false
		}

		seen |= rBit

	}
	return true
}
