package leetcodesgo

import (
	"fmt"
	"testing"
)

func TestParseBinaryNodesList(t *testing.T) {
	input := "[2,1,3,null,4,null,7]"
	root, err := parseBinaryNodeList(input)
	if err != nil {
		t.Errorf("Parsing error!")
	}
	fmt.Println(root, root.Left, root.Right)

	nodeValuesMap := map[*TreeNode]int{
		root: 2, root.Left: 1, root.Right: 3, root.Left.Right: 4, root.Right.Right: 7,
	}

	for node, value := range nodeValuesMap {
		if node.Val != value {
			t.Errorf("Wrong value! expected %d, got %d\n", value, node.Val)
		}
	}
}
