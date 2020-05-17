package subarray

// Time complexity: O(N)
// Space complexity: O(N)
func prefixSums(arr []int) []int {
	sums := make([]int, len(arr)+1)

	for i, a := range arr {
		sums[i+1] = sums[i] + a
	}

	return sums
}

// Time complexity: O(N)
// Space complexity: O(N)
func sumOfSubarray(arr []int, start, end int) int {
	sums := make([]int, len(arr)+1)
	for i, a := range arr {
		sums[i+1] = sums[i] + a
	}

	return sums[end+1] - sums[start]
}

// tags: prefix sums, sum of slice, sum of subarray
