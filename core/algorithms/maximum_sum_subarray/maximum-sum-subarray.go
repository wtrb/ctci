package subarray

// Find the maximum sum of the subarrays of arr.
// Time complexity: O(N)
// Space complexity: O(1)
func maxSubarraySum(arr []int) int {
	maxSum := arr[0]
	curSum := arr[0]

	for i := 1; i < len(arr); i++ {
		// curSum = max(arr[i]+curSum, arr[i])
		curSum = maxInts(curSum, 0) + arr[i]
		maxSum = maxInts(curSum, maxSum)
	}

	return maxSum
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// [-2, 2, 5, -11, 6]
// Max sum: 7

// tags: kadane, max slice, max sum subarray, dynamic programming
