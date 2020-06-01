package sort

// Time complexity: O(NlogN)
// Space complexity: O(1)
func HeapSort(arr []int) {
	buildMaxHeap(arr)

	for i := len(arr) - 1; i >= 1; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr[:i], 0)
	}
}

// Time complexity: O(N/2*logN)
func buildMaxHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i)
	}
}

// Time complexity: O(logN)
func heapify(arr []int, i int) {
	left := 2*i + 1
	right := 2*i + 2

	max := i
	if left < len(arr) && arr[left] > arr[max] {
		max = left
	}
	if right < len(arr) && arr[right] > arr[max] {
		max = right
	}

	if max != i {
		arr[i], arr[max] = arr[max], arr[i]
		heapify(arr, max)
	}
}

// tags: heap sort, max heap
