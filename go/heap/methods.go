package heap

func New[T any](cmp comparisonFunction[T]) *Heap[T] {
	return &Heap[T]{
		data:       make([]T, 0),
		comparison: cmp,
	}
}

func (heap *Heap[T]) Empty() bool {
	return len(heap.data) <= 0
}

func (heap *Heap[T]) Top() (T, error) {
	if heap.Empty() {
		var zero T
		return zero, ErrEmptyHeap
	}
	return heap.data[0], nil
}

// inserts element in heap, ordering as needed, then returns the current top element.
func (heap *Heap[T]) Push(item T) T {
	heap.data = append(heap.data, item)
	heap.orderHeap()
	return heap.data[0]
}

func (heap *Heap[T]) orderHeap() {
	index := len(heap.data) - 1
	for index > 0 && heap.comparison(heap.data[parentIndex(index)], heap.data[index]) {
		heap.data[index], heap.data[parentIndex(index)] = heap.data[parentIndex(index)], heap.data[index]
		index = parentIndex(index)
	}
}

func (heap *Heap[T]) Pop() (T, error) {
	if heap.Empty() {
		var zero T
		return zero, ErrEmptyHeap
	}
	top := heap.data[0]
	lastIndex := len(heap.data) - 1
	heap.data[0] = heap.data[lastIndex]
	heap.data = heap.data[:lastIndex]
	heap.orderHeap()

	return top, nil
}
