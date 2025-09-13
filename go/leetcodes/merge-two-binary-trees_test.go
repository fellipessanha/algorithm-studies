package leetcodesgo

import "testing"

func binaryTreeAreEqual(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	return t1 != nil && t2 != nil && t1.Val == t2.Val &&
		binaryTreeAreEqual(t1.Left, t2.Left) &&
		binaryTreeAreEqual(t1.Right, t2.Right)
}

func TestMergeTwoBinaryTrees(t *testing.T) {
	nodes1, _ := parseBinaryNodeList("[1,3,2,5]")
	nodes2, _ := parseBinaryNodeList("[2,1,3,null,4,null,7]")

	expectedResult, _ := parseBinaryNodeList("[3,4,5,5,4,null,7]")

	result := mergeTrees(nodes1, nodes2)

	if !binaryTreeAreEqual(expectedResult, result) {
		t.Error("Wrong result!")
	}
}
