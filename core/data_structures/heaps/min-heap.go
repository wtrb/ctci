package heap

import (
	"errors"
)

type MinHeap struct {
	items []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Peek() (int, error) {
	if h.Size() == 0 {
		return 0, errors.New("empty heap")
	}

	return h.items[0], nil
}

func (h *MinHeap) Poll() (int, error) {
	if h.Size() == 0 {
		return 0, errors.New("empty heap")
	}

	item := h.items[0]

	h.items[0] = h.items[h.Size()-1]
	h.items = h.items[:h.Size()-1]
	h.heapifyDown()

	return item, nil
}

func (h *MinHeap) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		smallerChildIdx := h.getLeftChildIndex(index)
		if h.hasRightChild(index) && h.rightChild(index) < h.leftChild(index) {
			smallerChildIdx = h.getRightChildIndex(index)
		}

		if h.items[index] < h.items[smallerChildIdx] {
			break

		} else {
			h.swap(index, smallerChildIdx)
		}

		index = smallerChildIdx
	}
}

func (h *MinHeap) Add(item int) {
	h.items = append(h.items, item)
	h.heapifyUp()
}

func (h *MinHeap) heapifyUp() {
	index := h.Size() - 1
	for h.hasParent(index) && h.parent(index) > h.items[index] {
		h.swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}
}

func (h *MinHeap) Size() int {
	return len(h.items)
}

func (h *MinHeap) ToSlice() []int {
	result := make([]int, h.Size())
	copy(result, h.items)
	return result
}

func (h *MinHeap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *MinHeap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *MinHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *MinHeap) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < h.Size()
}

func (h *MinHeap) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < h.Size()
}

func (h *MinHeap) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

func (h *MinHeap) leftChild(index int) int {
	return h.items[h.getLeftChildIndex(index)]
}

func (h *MinHeap) rightChild(index int) int {
	return h.items[h.getRightChildIndex(index)]
}

func (h *MinHeap) parent(index int) int {
	return h.items[h.getParentIndex(index)]
}

func (h *MinHeap) swap(idxOne, idxTwo int) {
	h.items[idxOne], h.items[idxTwo] = h.items[idxTwo], h.items[idxOne]
}

// tags: min heap
