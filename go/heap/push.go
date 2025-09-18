package heap

func New[T any](cmp comparisonFunction[T]) Heap[T] {
	return Heap[T]{
		data:       make([]T, 0),
		comparison: cmp,
	}
}
