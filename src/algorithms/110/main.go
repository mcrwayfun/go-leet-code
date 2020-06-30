package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/// Time Complexity: O(nlgN)
/// Space Complexity: O(h)
func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}

	if math.Abs(getHeight(root.Left) - getHeight(root.Right)) > 1{
		return false
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right){
		return false
	}

	return true
}

func getHeight(root *TreeNode) float64 {
	if root == nil{
		return 0
	}

	leftMaxDepth := getHeight(root.Left)
	rightMaxDepth := getHeight(root.Right)

	if leftMaxDepth > rightMaxDepth{
		return leftMaxDepth + 1
	}

	return rightMaxDepth + 1
}