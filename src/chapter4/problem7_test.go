package chapter4

import "testing"

func TestOrderProjects(t *testing.T) {

	projects := []string{"a", "b", "c", "d", "e", "f"}
	dependencies := [][]string{

		{"a", "d"},
		{"f", "b"},
		{"b", "d"},
		{"f", "a"},
		{"d", "c"},
	}

	ordered := OrderProjects(projects, dependencies)

	t.Logf("result: %v", ordered)

}

func TestOrderProjectsDFS(t *testing.T) {

	projects := []string{"a", "b", "c", "d", "e", "f"}
	dependencies := [][]string{

		{"a", "d"},
		{"f", "b"},
		{"b", "d"},
		{"f", "a"},
		{"d", "c"},
		//{"e", "d"},
	}

	ordered := OrderProjectsDFS(projects, dependencies)

	t.Logf("result: %v", ordered)

}
