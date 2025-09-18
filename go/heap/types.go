package heap

// should return true for l >= r, and false otherwise.
type comparisonFunction[T any] func(l, r T) bool

type Heap[T any] struct {
	data       []T
	comparison comparisonFunction[T]
}
