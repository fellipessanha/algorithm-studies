package leetcodesgo

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildBinaryTreeFromStringArray(values []string, idx int) (*TreeNode, error) {
	if idx >= len(values) || values[idx] == "null" {
		return nil, nil
	}

	value, err := strconv.Atoi(values[idx])
	if err != nil {
		return nil, err
	}
	left, leftErr := buildBinaryTreeFromStringArray(values, idx*2+1)
	if leftErr != nil {
		return nil, leftErr
	}
	right, rightErr := buildBinaryTreeFromStringArray(values, (idx+1)*2)
	if rightErr != nil {
		return nil, rightErr
	}
	return &TreeNode{value, left, right}, nil
}
