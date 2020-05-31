package sort

// InsertionSort sorts the input slice using insertion algorithm.
// Time complexity:
//		worst case (reverse sorted): O(n^2)
//		best case (already sorted): O(n)
//		average case: O(n^2)
// Space complexity: O(1)
// Usage: insertion sort runs in linear time on nearly sorted data.
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// tags: insertion sort
