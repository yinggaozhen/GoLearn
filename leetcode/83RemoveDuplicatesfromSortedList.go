package main

import (
	"fmt"
)

func main() {

	l1n3 := ListNode{
		Val:2,
	}

	l1n2 := ListNode{
		Val:1,
		Next:&l1n3,
	}

	l1n1 := ListNode{
		Val:1,
		Next:&l1n2,
	}

	fmt.Println(deleteDuplicates(&l1n1))
}

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	root := ListNode{
		Next: head,
	}

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}

		head = head.Next
	}

	return root.Next
}