package utils

import (
	"math"
)

type MinHeap[T any] struct {
	Heap       []T
	comparator func(T, T) bool
	isEqual    func(T, T) bool
}

func NewMinHeap[T any](comparator func(T, T) bool, isEqual func(T, T) bool) *MinHeap[T] {
	return &MinHeap[T]{comparator: comparator, isEqual: isEqual}
}

func (h *MinHeap[T]) Push(value T) {
	h.Heap = append(h.Heap, value)
	h.bubbleUp(len(h.Heap) - 1)
}

func (h *MinHeap[T]) Pop() T {
	last := len(h.Heap) - 1
	h.swap(0, last)
	value := h.Heap[last]
	h.Heap = h.Heap[:last]
	h.bubbleDown(0)
	return value
}

func (h *MinHeap[T]) Peek() T {
	return h.Heap[0]
}

func (h *MinHeap[T]) bubbleUp(index int) {
	parent := int(math.Floor(float64(index-1) / float64(2)))
	if parent < 0 || h.comparator(h.Heap[parent], h.Heap[index]) {
		return
	}
	h.swap(index, parent)
	h.bubbleUp(parent)
}

func (h *MinHeap[T]) bubbleDown(index int) {
	left := index*2 + 1
	right := index*2 + 2
	if left >= len(h.Heap) {
		return
	}
	smallest := left
	if right < len(h.Heap) && h.comparator(h.Heap[right], h.Heap[left]) {
		smallest = right
	}
	if h.comparator(h.Heap[index], h.Heap[smallest]) {
		return
	}
	h.swap(index, smallest)
	h.bubbleDown(smallest)
}

func (h *MinHeap[T]) swap(i int, j int) {
	h.Heap[i], h.Heap[j] = h.Heap[j], h.Heap[i]
}

func (h *MinHeap[T]) Len() int {
	return len(h.Heap)
}

func (h *MinHeap[T]) IsEmpty() bool {
	return h.Len() == 0
}

func (h *MinHeap[T]) Clear() {
	h.Heap = []T{}
}

func (h *MinHeap[T]) Contains(value T) bool {
	for _, item := range h.Heap {
		if h.isEqual(item, value) {
			return true
		}
	}
	return false
}
