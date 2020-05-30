package search

// Time complexity: O(logN)
// Space complexity: O(1)
func binarySearch(sortedArray []int, x int) bool {
	left := 0
	right := len(sortedArray) - 1

	for left <= right {
		mid := left + (right-left)/2

		if sortedArray[mid] == x {
			return true

		} else if x < sortedArray[mid] {
			right = mid - 1

		} else {
			left = mid + 1
		}
	}

	return false
}

// tags: binary search
