package sort

func MergeSort(arr []int) {
	temp := make([]int, len(arr))
	mergeSort(arr, temp, 0, len(arr)-1)
}

// Time complexity: O(NlogN)
// Space complexity: O(N)
func mergeSort(arr, temp []int, leftStart, rightEnd int) {
	if leftStart >= rightEnd {
		return
	}

	mid := (leftStart + rightEnd) / 2
	mergeSort(arr, temp, leftStart, mid)
	mergeSort(arr, temp, mid+1, rightEnd)

	mergeHalves(arr, temp, leftStart, rightEnd)
}

func mergeHalves(arr, temp []int, leftStart, rightEnd int) {
	leftEnd := (leftStart + rightEnd) / 2
	rightStart := leftEnd + 1

	left := leftStart
	right := rightStart
	index := leftStart
	for left <= leftEnd && right <= rightEnd {
		if arr[left] <= arr[right] {
			temp[index] = arr[left]
			left++

		} else {
			temp[index] = arr[right]
			right++
		}
		index++
	}
	// if left <= leftEnd {
	copy(temp[index:rightEnd+1], arr[left:leftEnd+1])
	// } else if right <= rightEnd {
	copy(temp[index:rightEnd+1], arr[right:rightEnd+1])
	// }

	copy(arr[leftStart:rightEnd+1], temp[leftStart:rightEnd+1])
}

// https://www.youtube.com/watch?v=KF2j-9iSf4Q

// tags: merge sort, devide and conquer
