package main

import (
	"fmt"
	"strconv"
)

func main() {
	ltree0103 := TreeNode{
		Val:3,
	}
	ltree0102 := TreeNode{
		Val:2,
	}
	ltree0101 := TreeNode{
		Val:1,
		Left: &ltree0102,
		Right: &ltree0103,
	}

	ltree0203 := TreeNode{
		Val:3,
	}
	ltree0202 := TreeNode{
		Val:2,
	}
	ltree0201 := TreeNode{
		Val:1,
		Left: &ltree0202,
		Right: &ltree0203,
	}

	fmt.Println(isSameTree(&ltree0101, &ltree0201))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Result100 struct {
	Val []string
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	lResult := Result100{}
	rResult := Result100{}

	parseTree(p, &lResult)
	parseTree(q, &rResult)

	if len(lResult.Val) == 0 || len(lResult.Val) != len(rResult.Val) {
		return false
	}

	for i := 0; i < len(lResult.Val); i++ {
		if rResult.Val[i] != lResult.Val[i] {
			return false
		}
	}

	return true
}

func parseTree(tree *TreeNode, result *Result100) {
	if tree != nil {
		result.Val = append(result.Val, strconv.Itoa(tree.Val))
		parseTree(tree.Left, result)
		parseTree(tree.Right, result)
	} else {
		result.Val = append(result.Val, "nil")
	}
}