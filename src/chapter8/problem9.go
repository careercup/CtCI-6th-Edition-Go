package chapter8

// Parens returns all valid parentheses by given number of pairs
func Parens(num int) []string {
	var (
		res []string
		bs  []byte
	)

	ParensR(num, 0, 0, bs, &res)

	return res
}

// ParensR adds valid parentheses with recursion
func ParensR(num, leftCnt, rightCnt int, bs []byte, res *[]string) {
	if num < leftCnt || num < rightCnt {
		return
	}
	if num == leftCnt && num == rightCnt {
		*res = append(*res, string(bs))
		return
	}

	if leftCnt < num {
		ParensR(num, leftCnt+1, rightCnt, append(bs, '('), res)
	}

	if rightCnt < num && rightCnt < leftCnt {
		ParensR(num, leftCnt, rightCnt+1, append(bs, ')'), res)
	}
}
