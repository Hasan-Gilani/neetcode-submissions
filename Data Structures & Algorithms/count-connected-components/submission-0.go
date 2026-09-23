func countComponents(n int, edges [][]int) int {
    graph := make(map[int][]int)
	visited := make(map[int]bool)
	count := 0

	for _, e := range edges {
		n1, n2 := e[0], e[1]
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}
	var dfs func(node int, graph map[int][]int, visited map[int]bool)
	dfs = func(node int, graph map[int][]int, visited map[int]bool) {
		visited[node] = true
		for _, c := range graph[node] {
			if _, ok := visited[c]; !ok {
				dfs(c, graph, visited)
			}
		}
	}
	for i := range n {
		if _, ok := visited[i]; !ok {
			dfs(i, graph, visited)
			count++
		}

	}
	return count
}
