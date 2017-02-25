package chapter8

// Permutations returns permutations of givien string
func Permutations(str string) []string {
	if len(str) == 0 {
		return []string{""}
	}

	var res []string
	for i := 0; i < len(str); i++ {
		subs := Permutations(TrimAt(str, i))
		for _, sub := range subs {
			res = append(res, string(str[i])+sub)
		}
	}

	return res
}

// TrimAt trims a char from given string by given index
func TrimAt(str string, i int) string {
	switch {
	case i < 0 || len(str) == 0 || i >= len(str):
		return str
	case i == len(str)-1:
		return str[:len(str)-1]
	case i == 0:
		return str[1:]
	default:
		return string(append([]byte(str[0:i]), str[i+1:]...))
	}
}
