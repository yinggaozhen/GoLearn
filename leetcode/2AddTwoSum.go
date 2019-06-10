package main

type ListNode2 struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	current := root
	add := 0

	for l1 != nil || l2 != nil || add != 0 {
		l1v, l2v := 0, 0
		if l1 != nil {
			l1v = l1.Val
		}

		if l2 != nil {
			l2v = l2.Val
		}

		currentVal := l1v + l2v + add
		add = currentVal / 10

		newNode := &ListNode{
			Val: currentVal % 10,
		}

		current.Next = newNode
		current = newNode

		if l1 != nil{
			l1 = l1.Next
		}

		if l2 != nil{
			l2 = l2.Next
		}
	}

	return root.Next
}