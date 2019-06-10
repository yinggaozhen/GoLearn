package main

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := int(math.Abs(float64(maxDepth110(root.Left)) - float64(maxDepth110(root.Right))))
	if diff >= 2 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth110(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return int(math.Max(float64(maxDepth110(node.Left)), float64(maxDepth110(node.Right)))) + 1
}