package heap

import (
	"math/rand"
	"testing"
)

func TestIndexesAreCorrect(t *testing.T) {
	correctValues := map[int][]int{
		0: {1, 2},
		1: {3, 4},
		2: {5, 6},
	}

	for key, values := range correctValues {
		if leftIndex(key) != values[0] {
			t.Errorf("Test failed! %d |> left should be %d, got %d\n", key, values[0], leftIndex(key))
		}
		if rightIndex(key) != values[1] {
			t.Errorf("Test failed! %d |> left should be %d, got %d\n", key, values[1], rightIndex(key))
		}
		if parentIndex(values[0]) != key {
			t.Errorf("Test failed! %d |> left should be %d, got %d\n", values[0], key, parentIndex(values[0]))
		}
		if parentIndex(values[1]) != key {
			t.Errorf("Test failed! %d |> left should be %d, got %d\n", values[1], key, parentIndex(values[1]))
		}
	}

}

func TestIndexesAreConsistent(t *testing.T) {
	testIndex := rand.Int() % (2 << 31)
	for i := 0; i < 10_000; i++ {
		rightTest := parentIndex(rightIndex(testIndex))
		if rightTest != testIndex {
			t.Errorf("Test failed! index %d |> right |> parent = %d\n", testIndex, rightTest)
		}
		leftTest := parentIndex(leftIndex(testIndex))
		if leftTest != testIndex {
			t.Errorf("Test failed! index %d |> left |> parent = %d\n", testIndex, leftTest)
		}
	}
}
