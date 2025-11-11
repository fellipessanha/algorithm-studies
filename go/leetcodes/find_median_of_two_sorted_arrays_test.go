package leetcodesgo

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	l1, l2 := []int{1, 4, 8}, []int{3, 4, 7}

	res := findMedianSortedArrays(l1, l2)
	if res != 4 {
		t.Errorf("Wrong value! expected %d", 4)
	}

	if findMedianSortedArrays(l1, append(l1, 9)) != 4 {
		t.Errorf("Wrong value! expected %d", 4)
	}
}
