package sort

// Time complexity:
// - worst case (reverse sorted): O(n^2)
// - best case (already sorted): O(n)
// - average case: O(n^2)
// Space complexity: O(1)
func BubbleSort(arr []int) {
	isSorted := false
	upper := len(arr) - 1
	for !isSorted {
		isSorted = true
		for i := 0; i < upper; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		upper--
	}
}

// tags: bubble sort
