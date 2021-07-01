package main

import (
	. "../tool/tree"
	"fmt"
)

type ResultType struct {
	IsBalanced bool
	MaxDepth   int
}

// time complexity: O(n)
// space complexity: O(1)
func isBalanced(root *TreeNode) bool {
	return helper(root).IsBalanced
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{true, 0}
	}

	leftResultType := helper(root.Left)
	rightResultType := helper(root.Right)

	if !leftResultType.IsBalanced || !rightResultType.IsBalanced {
		return ResultType{false, -1}
	}

	if abs(leftResultType.MaxDepth, rightResultType.MaxDepth) > 1 {
		return ResultType{false, -1}
	}

	return ResultType{true,
		max(leftResultType.MaxDepth, rightResultType.MaxDepth) + 1}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	nums := []int{3, 9, 20, NULL, NULL, 15, 7}
	root := Ints2TreeNode(nums)
	fmt.Println(isBalanced(root))

	nums = []int{1, 2, 2, 3, 3, NULL, NULL, 4, 4}
	root = Ints2TreeNode(nums)
	fmt.Println(isBalanced(root))
}
