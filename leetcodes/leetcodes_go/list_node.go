package leetcodesgo

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// String implements the fmt.Stringer interface for ListNode
// This allows fmt.Println(node) to display the node as an array
func (node *ListNode) String() string {
	if node == nil {
		return "[]"
	}

	var asString string
	curr := node
	for curr != nil {
		asString += " " + strconv.Itoa(curr.Val)
		curr = curr.Next
	}

	return "[" + strings.Trim(asString, " ") + "]"
}

func buildListNode(values []int) (root *ListNode) {
	if len(values) == 0 {
		return nil
	}
	root = &ListNode{Val: values[0]}
	node := root
	for _, val := range values[1:] {
		node.Next = &ListNode{Val: val}
		node = node.Next
	}
	return
}
