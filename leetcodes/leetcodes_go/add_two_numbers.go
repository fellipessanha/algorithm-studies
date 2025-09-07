package leetcodesgo

func getAsNumber(node *ListNode) (ret int) {
	decimal_multiplier := 1
	for node != nil {
		ret = ret + node.Val*decimal_multiplier
		node = node.Next
		decimal_multiplier *= 10
	}
	return
}

func iterateNode(node *ListNode) (int, *ListNode) {
	if node == nil {
		return 0, nil
	}
	return node.Val, node.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum_root := &ListNode{0, nil}
	node := sum_root
	plusOne := false

	for l1 != nil || l2 != nil {
		v1, n1 := iterateNode(l1)
		l1 = n1
		v2, n2 := iterateNode(l2)
		l2 = n2

		current := v1 + v2
		if plusOne {
			current++
		}

		plusOne = current > 9
		node.Next = &ListNode{current % 10, nil}
		node = node.Next
	}
	if plusOne {
		node.Next = &ListNode{1, nil}
	}

	return sum_root.Next
}
