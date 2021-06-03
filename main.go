package main

import (
	"fmt"

	maxheap "github.com/wtrb/ctci/core/data_structures/heaps/max_heap"
	minheap "github.com/wtrb/ctci/core/data_structures/heaps/min_heap"
)

func main() {
	min := minheap.New()
	min.Add(1)
	min.Add(5)
	min.Add(4)
	min.Add(8)
	min.Add(9)
	min.Add(7)
	min.Add(10)
	min.Add(0)

	minSlice := min.ToSlice()
	fmt.Println(minSlice)

	max := maxheap.New()
	max.Add(1)
	max.Add(5)
	max.Add(4)
	max.Add(8)
	max.Add(9)
	max.Add(7)
	max.Add(10)
	max.Add(0)

	maxSlice := max.ToSlice()
	fmt.Println(maxSlice)
}
