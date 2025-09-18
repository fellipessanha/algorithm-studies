package leetcodesgo

import (
	"testing"
)

func TestParseBinaryNodesList(t *testing.T) {
	input := "[2,1,3,null,4,null,7]"
	root, err := parseBinaryNodeList(input)
	if err != nil {
		t.Errorf("Parsing error!")
	}

	nodeValuesMap := map[*TreeNode]int{
		root: 2, root.Left: 1, root.Right: 3, root.Left.Right: 4, root.Right.Right: 7,
	}

	for node, value := range nodeValuesMap {
		if node.Val != value {
			t.Errorf("Wrong value! expected %d, got %d\n", value, node.Val)
		}
	}
}

func TestParseMixedArrayArray(t *testing.T){
	input := "[[\"kimchi\",\"miso\",\"sushi\",\"moussaka\",\"ramen\",\"bulgogi\"],[\"korean\",\"japanese\",\"japanese\",\"greek\",\"japanese\",\"korean\"],[9,12,8,15,14,7]]"
	output, err := parseStringArrayArray(input)
	if err != nil {
		t.Errorf("Parsing error!")
	}

	if len(output) != 3{
		t.Errorf("Parsing error! wrong number of arrays")
	}

	for _, arr := range output {
		if len(arr) != 6 {
			t.Errorf("Parsing error! wrong number of elements in array")
		}
	}

}