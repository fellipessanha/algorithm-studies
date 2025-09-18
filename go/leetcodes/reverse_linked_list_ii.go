package leetcodesgo

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	nodesList := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		nodesList = append(nodesList, node.Val)
	}

	for i := 0; i <= (right-left)/2; i++ {
		l := left - 1 + i
		r := right - 1 - i
		nodesList[l], nodesList[r] = nodesList[r], nodesList[l]
	}

	retHead := &ListNode{Val: -1}
	retNode := retHead
	for _, value := range nodesList {
		retNode.Next = &ListNode{Val: value}
		retNode = retNode.Next
	}

	return retHead.Next
}
