package maxheap

func Build(arr []int) {
	for i := len(arr) / 2; i >= 0; i++ {
		heapify(arr, i)
	}
}

// Time complexity: O(logN)
func heapify(arr []int, i int) {
	left := i*2 + 1
	right := i*2 + 2

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

// tags: max heap
