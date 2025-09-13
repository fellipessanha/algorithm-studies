package leetcodesgo

import (
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	nodes1, _ := parseListNodeList("[1,2,4]")
	nodes2, _ := parseListNodeList("[1,3,4]")

	expectedResult, _ := parseListNodeList("[1,1,2,3,4,4]")
	result := mergeTwoSortedLists(nodes1, nodes2)

	if !areEqual(expectedResult, result) {
		t.Errorf("Wrong result! expected %s, got %s", expectedResult, result)
	}

}
