package chapter4

type Project struct {
	Name string

	Children []*Project
	DepCnt   int
}

type ProjectNode struct {
	Name string

	State int // 0 new 1 checking 2 done

	Dependencies []*ProjectNode
}

func OrderProjects(projects []string, Dependencies [][]string) []string {
	projectMap := make(map[string]*Project, len(projects))
	for _, name := range projects {
		projectMap[name] = &Project{Name: name, Children: make([]*Project, 0), DepCnt: 0}
	}
	for _, dep := range Dependencies {
		from, to := projectMap[dep[0]], projectMap[dep[1]]
		to.DepCnt++
		from.Children = append(from.Children, to)

	}

	ordered := make([]string, len(projects))
	index := 0
	for len(projectMap) > 0 {
		for _, project := range projectMap {
			if project.DepCnt == 0 {
				ordered[index] = project.Name
				index++
				for _, child := range project.Children {
					child.DepCnt--
				}
				delete(projectMap, project.Name)
			}
		}
	}
	return ordered
}

func OrderProjectsDFS(projects []string, dependencies [][]string) []string {
	projectNodes := make(map[string]*ProjectNode, len(projects))

	for _, name := range projects {
		projectNodes[name] = &ProjectNode{Name: name, Dependencies: make([]*ProjectNode, 0), State: 0}
	}
	for _, dep := range dependencies {
		from, to := projectNodes[dep[0]], projectNodes[dep[1]]

		to.Dependencies = append(to.Dependencies, from)

	}

	ordered := make([]string, len(projects))
	index := 0

	for _, project := range projects {
		var projectNode *ProjectNode
		found := false
		if projectNode, found = projectNodes[project]; !found {
			ordered[index] = project
			index++
			continue
		}
		if !dfs(projectNode, &ordered, &index) {
			return nil
		}
	}

	return ordered
}

func dfs(projectNode *ProjectNode, ordered *[]string, index *int) bool {
	if projectNode.State == 2 {
		return true
	}
	if projectNode.State == 1 {
		return false
	}
	projectNode.State = 1
	for _, node := range projectNode.Dependencies {

		if ok := dfs(node, ordered, index); !ok {
			return false
		}
	}
	(*ordered)[*index] = projectNode.Name
	*index++
	projectNode.State = 2
	return true

}
