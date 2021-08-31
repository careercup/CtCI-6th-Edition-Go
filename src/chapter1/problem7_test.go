package chapter1

import (
	"reflect"
	"testing"
)

func TestMatrixRotate(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			[][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}
	for _, c := range cases {
		MatrixRotate(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Fatalf("Expected: %v, actual: %v\n", c.expected, c.input)
		}
	}
}
