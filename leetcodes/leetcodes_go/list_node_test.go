package leetcodesgo

import (
	"fmt"
	"testing"
)

func TestNodeGenerator(t *testing.T) {
	array := []int{1, 2, 3}
	node := buildListNode(array)

	for _, value := range array {
		if node == nil {
			t.Errorf("Nodes list smaller than supplied array")
		}
		if node.Val != value {
			t.Errorf("Node expected value %d, got %d", value, node.Val)
		}
		node = node.Next
	}

}

func TestListNodeAsString(t *testing.T) {
	array := []int{4, 5, 6}
	nodes := buildListNode(array)

	arrayString := fmt.Sprintln(array)
	nodesString := fmt.Sprintln(nodes)

	if arrayString != nodesString {
		t.Errorf("array's output is different from generator Node. expected %s, got %s", arrayString, nodesString)
	}

}
