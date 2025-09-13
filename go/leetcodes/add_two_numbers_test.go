package leetcodesgo

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var l1, l2 = []int{2, 4, 3}, []int{5, 6, 4}

	n := getAsNumber(addTwoNumbers(buildListNode(l1), buildListNode(l2)))
	expectedResult := 807

	if n != expectedResult {
		t.Errorf("add_two_numbers solution is wrong! expected %d, got %d", expectedResult, n)
	}

}
