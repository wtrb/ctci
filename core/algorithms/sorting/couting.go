package sort

// Time complexity: (N+K) with N is input array size and K is additional bucket size.
// Space complexity: O(K)
func CountingSort(arr []int) {
	min := 1<<63 - 1
	max := -1 << 63
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	bucket := make([]int, max-min+1)
	for _, v := range arr {
		bucket[v-min]++
	}

	idx := 0
	for i, v := range bucket {
		if v > 0 {
			for j := 0; j < v; j++ {
				arr[idx] = i + min
				idx++
			}
		}
	}
}

// tags: counting sort
