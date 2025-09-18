package leetcodesgo

import "testing"

func TestReverseLinkedListII(t *testing.T) {
	tests := map[string][]int{"[3,5]": {1, 2, 5, 3}, "[1,2,3,4,5]": {2, 4, 1, 4, 3, 2, 5}, "[1,2,3]": {1, 3, 3, 2, 1}, "[1]": {1, 1, 1}}
	for input, positions := range tests {
		nodesList, parseErr := parseListNodeList(input)
		if parseErr != nil {
			t.Errorf("Parsing error!")
		}
		output := reverseBetween(nodesList, positions[0], positions[1])
		expectedList := buildListNode(positions[2:])
		for output != nil && expectedList != nil {
			if output.Val != expectedList.Val {
				t.Errorf("reverse_linked_list_ii solution is wrong! expected %d, got %d", expectedList.Val, output.Val)
			}
			output = output.Next
			expectedList = expectedList.Next
		}
		if output != nil || expectedList != nil {
			t.Errorf("reverse_linked_list_ii solution is wrong! lists have different lengths")
		}

	}
}
