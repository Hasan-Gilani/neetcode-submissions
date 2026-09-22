/* (0, 2)
Input: [
  [2147483647,	-1,				0,					2147483647],
  [2147483647,	2147483647,	2147483647,						-1],
  [2147483647, 	-1,			2147483647,						-1],
  [0,			-1,			2147483647,				2147483647]

	[[2147483647,-1,0,2147483648],
	[2,2147483648,2147483648,-1],
	[1,-1,2,-1],
	[0,-1,3,2147483648]]

] */

var INF = 2147483647

func islandsAndTreasure(grid [][]int) {
	queue := [][]int{}
	for i := 0; i < len(grid); i++{
		for j:=0 ; j < len(grid[i]); j++{
			if grid[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	for len(queue) != 0 {
		x := queue[0]
		i, j := x[0], x[1]
		if (i + 1) < len(grid) && grid[i + 1][j] == INF {
			grid[i + 1][j] = grid[i][j] + 1
			queue = append(queue, []int{i+1, j})
		}
		if (i - 1) >= 0 && grid[i - 1][j] == INF {
			grid[i -1][j] = grid[i][j] + 1
			queue = append(queue, []int{i-1, j})
		}
		if (j + 1) < len(grid[i]) && grid[i][j + 1] == INF {
			grid[i][j + 1] = grid[i][j] + 1
			queue = append(queue, []int{i, j + 1})
		}
		if (j - 1) >= 0 && grid[i][j - 1] == INF {
			grid[i][j-1] = grid[i][j] + 1
			queue = append(queue, []int{i, j - 1})
		}
		queue = queue[1:len(queue)]
	}
}
