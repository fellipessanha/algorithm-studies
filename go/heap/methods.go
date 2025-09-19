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
	heap.orderHeapUp()
	return heap.data[0]
}

func (heap *Heap[T]) orderHeapUp() {
	index := len(heap.data) - 1
	for index > 0 && heap.comparison(heap.data[index], heap.data[parentIndex(index)]) {
		heap.data[index], heap.data[parentIndex(index)] = heap.data[parentIndex(index)], heap.data[index]
		index = parentIndex(index)
	}
}

func (heap *Heap[T]) orderHeapDown() {
	for index := 0; leftIndex(index) < len(heap.data); {
		l, r := childrenIndexes(index)
		if r >= len(heap.data) {
			r = l
		}
		if heap.comparison(heap.data[index], heap.data[l]) && heap.comparison(heap.data[index], heap.data[r]) {
			break
		}
		newIndex := l
		if heap.comparison(heap.data[r], heap.data[l]) {
			newIndex = r
		}
		heap.data[index], heap.data[newIndex] = heap.data[newIndex], heap.data[index]
		index = newIndex
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
	heap.orderHeapDown()

	return top, nil
}
