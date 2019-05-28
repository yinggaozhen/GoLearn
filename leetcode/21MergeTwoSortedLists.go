package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1n3 := ListNode{
		Val:4,
	}

	l1n2 := ListNode{
		Val:2,
		Next:&l1n3,
	}

	l1n1 := ListNode{
		Val:1,
		Next:&l1n2,
	}

	l2n3 := ListNode{
		Val:4,
	}

	l2n2 := ListNode{
		Val:3,
		Next:&l2n3,
	}

	l2n1 := ListNode{
		Val:1,
		Next:&l2n2,
	}

	result := mergeTwoLists(&l1n1, &l2n1)
	fmt.Println(result)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}
	current := &result

	for l1 != nil || l2 != nil {

		node := ListNode{}

		if l1 == nil {
			node.Val = l2.Val
			current.Next = &node
			l2 = l2.Next
		} else if l2 == nil {
			node.Val = l1.Val
			current.Next = &node
			l1 = l1.Next
		} else if l1.Val <= l2.Val {
			node.Val = l1.Val
			current.Next = &node
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			current.Next = &node
			l2 = l2.Next
		}

		current = &node
	}

	return result.Next
}

//func forward(node ListNode, current *ListNode, list *ListNode) {
//	node.Val = list.Val
//	current.Next = &node
//	list = list.Next
//}