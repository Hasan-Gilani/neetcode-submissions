func orangesRotting(grid [][]int) int {
	minutes := 0
    rottenFruits := [][]int{}
	for i := 0; i < len(grid); i++{
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				rottenFruits = append(rottenFruits, []int{i, j, 0})
			}
		}
	}
	for len(rottenFruits) != 0 {
		f := rottenFruits[0]
		i, j, t := f[0], f[1], f[2]

		if i+1 < len(grid) && grid[i+1][j] == 1 {
			grid[i+1][j] = 2
			rottenFruits = append(rottenFruits, []int{i+1, j, t + 1})
			minutes = t + 1
		}
		if i > 0 && grid[i-1][j] == 1 {
			grid[i-1][j] = 2
			rottenFruits = append(rottenFruits, []int{i-1, j, t + 1})
			minutes = t + 1
		}
		if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
			grid[i][j+1] = 2
			rottenFruits = append(rottenFruits, []int{i, j+1, t + 1})
			minutes = t + 1
		}
		if j > 0 && grid[i][j-1] == 1 {
			grid[i][j-1] = 2
			rottenFruits = append(rottenFruits, []int{i, j-1, t + 1})
			minutes = t + 1
		}

		rottenFruits = rottenFruits[1:len(rottenFruits)]
	}

	for i := 0; i < len(grid); i++{
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return minutes
}
