package flood

func numIslands(grid [][]byte) int {
	num := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				num++
				dfs(grid, i, j)
			}
		}
	}

	return num
}

// sink the entire island
func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}