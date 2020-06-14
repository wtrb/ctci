package wines

func MaxProfit(prices []int) int {
	// return maxProfit(prices, 1, 0, len(prices)-1)
	// return maxProfit(prices, 0, len(prices)-1)

	cache := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		cache[i] = make([]int, len(prices))
	}
	return maxProfit(prices, cache, 0, len(prices)-1)
}

// step 1: Write a backtrack
// Time complexity: O(2^N)
// func maxProfit(prices []int, year int, left, right int) int {
// 	if left > right {
// 		return 0
// 	}

// 	return maxInts(
// 		maxProfit(prices, year+1, left+1, right)+year*prices[left],
// 		maxProfit(prices, year+1, left, right-1)+year*prices[right],
// 	)
// }

// step 2: Minimize the state space of function arguments
// Time complexity: O(2^N)
// func maxProfit(prices []int, left, right int) int {
// 	if left > right {
// 		return 0
// 	}

// 	year := len(prices) - (right - left + 1) + 1

// 	return maxInts(
// 		maxProfit(prices, left+1, right)+year*prices[left],
// 		maxProfit(prices, left, right-1)+year*prices[right],
// 	)
// }

// step 3: cache
// Time complexity: O(N^2)
func maxProfit(prices []int, cache [][]int, left, right int) int {
	if left > right {
		return 0
	}

	if cache[left][right] != 0 {
		return cache[left][right]
	}

	year := len(prices) - (right - left + 1) + 1
	cache[left][right] = maxInts(
		maxProfit(prices, cache, left+1, right)+year*prices[left],
		maxProfit(prices, cache, left, right-1)+year*prices[right],
	)

	return cache[left][right]
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://www.quora.com/Are-there-any-good-resources-or-tutorials-for-dynamic-programming-DP-besides-the-TopCoder-tutorial/answer/Michal-Danil%C3%A1k?share=1&srid=3OBi

// tags: dp, dynamic programming
