package subarray

// Time complexity: O(N)
// Space complexity: O(N)
func prefixSums(arr []int) []int {
	sums := make([]int, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		sums[i+1] = arr[i] + sums[i]
	}

	return sums
}

// Time complexity: O(N)
// Space complexity: O(N)
func sumOfSubarray(arr []int, start, end int) int {
	prefixSums := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		prefixSums[i+1] = arr[i] + prefixSums[i]
	}

	return prefixSums[end+1] - prefixSums[start]
}

// tags: prefix sums, sum of slice, sum of subarray, dp, dynamic programming
