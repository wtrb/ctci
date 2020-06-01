package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

// Time complexity:
//		average - O(NlogN)
//		worst - O(N^2)
// Space complexity: O(1)
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivot := arr[(left+right)/2]
	index := partition(arr, left, right, pivot)
	quickSort(arr, left, index-1)
	quickSort(arr, index, right)
}

func partition(arr []int, left, right int, pivot int) int {
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}

// https://www.youtube.com/watch?v=SLauY6PpjW4

// tags: quick sort, devide and conquer
