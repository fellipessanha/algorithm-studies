package heap

import (
	"testing"
)

func TestHeapMethods(t *testing.T) {
	// Create a new max-heap for integers
	maxHeap := New(func(p, c int) bool { return p > c })

	// Test if the heap is initially empty
	if !maxHeap.Empty() {
		t.Error("Expected heap to be empty")
	}

	// Test error when accessing top of empty heap
	_, err := maxHeap.Top()
	if err == nil {
		t.Error("Expected error when accessing top of empty heap")
	}
	pushOrder := []struct {
		insert   int
		expected int
	}{{5, 5}, {3, 5}, {8, 8}, {1, 8}}

	// Test pushing elements into the heap
	for _, pushOrder := range pushOrder {
		cur, max := pushOrder.insert, pushOrder.expected
		maxHeap.Push(cur)
		currentTop, pushErr := maxHeap.Top()
		if pushErr != nil {
			t.Errorf("Unexpected error when getting top element: %v", pushErr)
		}
		if currentTop != max {
			t.Errorf("Expected top to be %d after pushing %d, got %d", max, cur, currentTop)
		}
	}

	// Test if the heap is not empty after pushes
	if maxHeap.Empty() {
		t.Error("Expected heap to be non-empty")
	}

	// Test popping elements from the heap
	for _, expectedTop := range []int{8, 5, 3, 1} {
		poppedValue, popErr := maxHeap.Pop()
		if popErr != nil {
			t.Errorf("Unexpected error when popping element: %v", popErr)
		}
		if poppedValue != expectedTop {
			t.Errorf("Expected popped value to be %d, got %d", expectedTop, poppedValue)
		}
	}
}
