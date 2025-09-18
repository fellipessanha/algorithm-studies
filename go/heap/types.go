package heap

import "errors"

// should return true for l >= r, and false otherwise.
type comparisonFunction[T any] func(parent, children T) bool

type Heap[T any] struct {
	data       []T
	comparison comparisonFunction[T]
}

var ErrEmptyHeap = errors.New("Heap is empty")
