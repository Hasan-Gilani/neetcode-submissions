// [0, 1] 1 -> 0

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for i := 0; i < len(prerequisites); i++ {
		preq := prerequisites[i]
		course, prereq := preq[0], preq[1]
		graph[prereq] = append(graph[prereq], course)
		inDegree[course]++
	}

	queue := []int{}
	for c := 0; c < numCourses; c++ {
		if inDegree[c] == 0 {
			queue = append(queue, c)
		}
	}

	taken := 0
	for len(queue) != 0 {
		taken++
		preq := queue[0]
		queue = queue[1:len(queue)]

		for _, c := range graph[preq] {
			inDegree[c]--
			if inDegree[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	return taken == numCourses
}
