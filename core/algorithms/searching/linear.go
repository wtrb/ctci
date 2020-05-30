package search

// Time complexity: O(N)
// Space complexity: O(1)
func linearSearch(array []int, x int) bool {
	for _, v := range array {
		if v == x {
			return true
		}
	}
	return false
}

// tags: linear search
