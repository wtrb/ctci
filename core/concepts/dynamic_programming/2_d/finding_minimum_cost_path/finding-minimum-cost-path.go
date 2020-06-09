package costpath

// Time complexity: O(N*M)
// Space complexity: O(N*M)
func minCost(cost [][]int) int {
	rows := len(cost)
	cols := len(cost[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	dp[0][0] = cost[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + cost[i][0]
	}
	for j := 1; j < cols; j++ {
		dp[0][j] = dp[0][j-1] + cost[0][j]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = minInts(dp[i-1][j], dp[i][j-1]) + cost[i][j]
		}
	}

	return dp[rows-1][cols-1]
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinCost(i,j) = min(MinCost(i-1,j),MinCost(i,j-1)) + Cost[i][j]

// tags: dp, dynamic programming