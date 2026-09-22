func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make(map[int][]int)
	inDegree := make(map[int]int)
	order := []int{}

	for _, p := range prerequisites {
		pre, course := p[1], p[0]
		graph[pre] = append(graph[pre], course)
		inDegree[course]++
	}
	queue := []int{}
	for c := 0; c < numCourses; c++ {
		if inDegree[c] == 0 {
			queue = append(queue, c)
		}
	}
	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:len(queue)]
		
		order = append(order, c)
		for _, course := range graph[c] {
			inDegree[course]--
			if inDegree[course] == 0 {
				queue = append(queue, course)
			}
		}
	}
	if len(order) < numCourses {
		return []int{}
	}
	return order
}
