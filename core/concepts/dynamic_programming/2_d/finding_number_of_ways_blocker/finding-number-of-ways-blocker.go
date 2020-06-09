package path

const BLOCK = 1

func ways(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	for i := 1; i < rows; i++ {
		if matrix[i][0] == BLOCK {
			break
		}
		dp[i][0] = 1
	}
	for j := 1; j < cols; j++ {
		if matrix[0][j] == BLOCK {
			break
		}
		dp[0][j] = 1
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] != BLOCK {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[rows-1][cols-1]
}
