package leetcodesgo

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	entryArray := []int{1, 3, 5, 6}
	searchEntries := map[int]int{0: 0, 5: 2, 2: 1, 8: 4}
	for target, expectedResult := range searchEntries {
		result := searchInsert(entryArray, target)
		if expectedResult != result {
			t.Errorf("Wrong Result! expected %d, got %d", expectedResult, result)
		}
	}
}
