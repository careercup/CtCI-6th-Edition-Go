package chapter8

import (
	"reflect"
	"testing"
)

func TestPowerSet(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out [][]int
	}{

		"Should return empty subset when input is empty": {
			in: []int{},
			out: [][]int{
				[]int{},
			},
		},

		"Should return 2 subsets when input's length is 1": {
			in: []int{1},
			out: [][]int{
				[]int{},
				[]int{1},
			},
		},

		"Should return 4 subsets when input's length is 2": {
			in: []int{1, 2},
			out: [][]int{
				[]int{},
				[]int{1},
				[]int{2},
				[]int{2, 1},
			},
		},

		"Should return 8 subsets when input's length is 3": {
			in: []int{1, 2, 3},
			out: [][]int{
				[]int{},
				[]int{1},
				[]int{2},
				[]int{2, 1},
				[]int{3},
				[]int{3, 1},
				[]int{3, 2},
				[]int{3, 2, 1},
			},
		},

		"Should return 16 subsets when input's length is 4": {
			in: []int{1, 2, 3, 4},
			out: [][]int{
				[]int{},
				[]int{1},
				[]int{2},
				[]int{2, 1},
				[]int{3},
				[]int{3, 1},
				[]int{3, 2},
				[]int{3, 2, 1},
				[]int{4},
				[]int{4, 1},
				[]int{4, 2},
				[]int{4, 2, 1},
				[]int{4, 3},
				[]int{4, 3, 1},
				[]int{4, 3, 2},
				[]int{4, 3, 2, 1},
			},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := PowerSet(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%#v expected=%#v", out, test.out)
			}
		})
	}
}
