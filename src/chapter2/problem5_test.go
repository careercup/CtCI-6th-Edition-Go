package chapter2

import (
	"reflect"
	"testing"
)

func TestAddTwoLists(t *testing.T) {
	cases := []struct {
		vals1, vals2 []int
		expected     int
	}{
		{
			[]int{7, 1, 6},
			[]int{5, 9, 2},
			912,
		},
		{
			[]int{7, 1, 6},
			[]int{5, 9, 8},
			1512,
		},
		{
			[]int{7, 1, 6},
			[]int{2, 4},
			659,
		},
		{
			[]int{2, 4},
			[]int{7, 1, 6},
			659,
		},
		{
			[]int{},
			[]int{7, 1, 6},
			617,
		},
		{
			[]int{7, 1, 6},
			[]int{},
			617,
		},
	}
	for _, c := range cases {
		l1 := GetLinkedListFromValues(c.vals1)
		l2 := GetLinkedListFromValues(c.vals2)
		actual := AddTwoLists(l1, l2)
		if c.expected != actual {
			t.Fatalf("Expected: %v, actual: %v\n", c.expected, actual)
		}
	}
}

func TestAddTwoListsB(t *testing.T) {
	cases := []struct {
		vals1, vals2, expected []int
	}{
		{
			[]int{7, 1, 6},
			[]int{5, 9, 2},
			[]int{2, 1, 9},
		},
		{
			[]int{7, 1, 6},
			[]int{5, 9, 8},
			[]int{2, 1, 5, 1},
		},
		{
			[]int{7, 1, 6},
			[]int{2, 4},
			[]int{9, 5, 6},
		},
		{
			[]int{2, 4},
			[]int{7, 1, 6},
			[]int{9, 5, 6},
		},
		{
			[]int{},
			[]int{7, 1, 6},
			[]int{7, 1, 6},
		},
		{
			[]int{7, 1, 6},
			[]int{},
			[]int{7, 1, 6},
		},
	}
	for _, c := range cases {
		l1 := GetLinkedListFromValues(c.vals1)
		l2 := GetLinkedListFromValues(c.vals2)

		actual := AddTwoListsB(l1, l2).Slice()

		if !reflect.DeepEqual(c.expected, actual) {
			t.Fatalf("Expected: %v, actual: %v\n", c.expected, actual)
		}
	}
}
