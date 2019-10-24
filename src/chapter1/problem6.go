package chapter1

import (
	"bytes"
	"strconv"
)

func BasicCompress(input string) string {
	var compStr string
	counter := 1
	for i := 1; i < len(input); i++ {
		if input[i-1] == input[i] {
			counter++
		} else {
			compStr = compStr + string(input[i-1]) + strconv.Itoa(counter)
			counter = 1
		}
	}
	compStr = compStr + string(input[len(input)-1]) + strconv.Itoa(counter)

	if len(compStr) > len(input) {
		return input
	} else {
		return compStr
	}
}

func BasicCompressB(input string) string {
	var compressed bytes.Buffer
	counter := 1
	inputLen := len(input)
	for i := 1; i < inputLen; i++ {

		if input[i-1] == input[i] {
			counter++
		} else {
			compressed.WriteByte(input[i-1])
			compressed.WriteString(strconv.Itoa(counter))
			counter = 1
		}
		if compressed.Len() >= inputLen {
			return input
		}
	}
	compressed.WriteByte(input[inputLen-1])
	compressed.WriteString(strconv.Itoa(counter))
	if compressed.Len() >= inputLen {
		return input
	}
	return compressed.String()
}
