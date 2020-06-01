package maxheap

import (
	"errors"
)

type MaxHeap struct {
	items []int
}

func New() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Peek() (int, error) {
	if h.Size() == 0 {
		return 0, errors.New("empty heap")
	}

	return h.items[0], nil
}

func (h *MaxHeap) Poll() (int, error) {
	if h.Size() == 0 {
		return 0, errors.New("empty heap")
	}

	item := h.items[0]

	h.items[0] = h.items[h.Size()-1]
	h.items = h.items[:h.Size()-1]
	h.heapifyDown()

	return item, nil
}

func (h *MaxHeap) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		greaterChildIdx := h.getLeftChildIndex(index)
		if h.hasRightChild(index) && h.rightChild(index) > h.leftChild(index) {
			greaterChildIdx = h.getRightChildIndex(index)
		}

		if h.items[index] > h.items[greaterChildIdx] {
			break

		} else {
			h.swap(index, greaterChildIdx)
		}

		index = greaterChildIdx
	}
}

func (h *MaxHeap) Add(item int) {
	h.items = append(h.items, item)
	h.heapifyUp()
}

func (h *MaxHeap) heapifyUp() {
	index := h.Size() - 1
	for h.hasParent(index) && h.parent(index) < h.items[index] {
		h.swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}
}

func (h *MaxHeap) Size() int {
	return len(h.items)
}

func (h *MaxHeap) ToSlice() []int {
	result := make([]int, h.Size())
	copy(result, h.items)
	return result
}

func (h *MaxHeap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *MaxHeap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *MaxHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *MaxHeap) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < h.Size()
}

func (h *MaxHeap) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < h.Size()
}

func (h *MaxHeap) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

func (h *MaxHeap) leftChild(index int) int {
	return h.items[h.getLeftChildIndex(index)]
}

func (h *MaxHeap) rightChild(index int) int {
	return h.items[h.getRightChildIndex(index)]
}

func (h *MaxHeap) parent(index int) int {
	return h.items[h.getParentIndex(index)]
}

func (h *MaxHeap) swap(idxOne, idxTwo int) {
	h.items[idxOne], h.items[idxTwo] = h.items[idxTwo], h.items[idxOne]
}

// tags: min heap
