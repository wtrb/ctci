package bitmnp

import (
	"math"
)

// Time complexity: O(2^n * n) with n is set size.
// Space complexity: O(1)
func countSubsets(set []int, val int) int {
	n := len(set)
	count := 0

	for x := 0; x < int(math.Pow(2, float64(n))); x++ { // go over all possible subsets 2^n - 1
		sum := 0

		for i := uint(0); i < uint(n); i++ {
			if x&(1<<i) != 0 {
				sum += set[i]
			}
		}

		if sum >= val {
			count++
		}
	}

	return count
}

// count how many subsets have sum of elements greater than or equal to a given value.
