package heap

import "testing"

func TestHeapMethods(t *testing.T) {
	// Create a new min-heap for integers
	minHeap := New(func(p, c int) bool { return p > c })

	// Test if the heap is initially empty
	if !minHeap.Empty() {
		t.Error("Expected heap to be empty")
	}

	// Test error when accessing top of empty heap
	_, err := minHeap.Top()
	if err == nil {
		t.Error("Expected error when accessing top of empty heap")
	}

	// Test pushing elements into the heap
	minHeap.Push(5)
	minHeap.Push(3)
	minHeap.Push(8)
	minHeap.Push(1)

	// Test if the heap is not empty after pushes
	if minHeap.Empty() {
		t.Error("Expected heap to be non-empty")
	}

	// Test the top element of the heap
	top, err := minHeap.Top()
	if err != nil {
		t.Errorf("Unexpected error when getting top element: %v", err)
	}
	if top != 1 {
		t.Errorf("Expected top element to be 1, got %d", top)
	}

	// Test pushing another element and checking the top again
	minHeap.Push(0)

	// Test popping elements from the heap
	expectedOrder := []int{0, 1, 3, 5, 8}
	for _, expected := range expectedOrder {
		popped, err := minHeap.Pop()
		if err != nil {
			t.Errorf("Unexpected error when popping element: %v", err)
		}
		if popped != expected {
			t.Errorf("Expected popped element to be %d, got %d", expected, popped)
		}
	}
}
