package main

import (
	. "../tool/tree"
	"fmt"
)

// time complexity: O(n)
// space complexity: O(1)
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var result int
	dfs(root, low, high, &result)
	return result
}

func dfs(root *TreeNode, low int, high int, result *int) int {
	if root == nil {
		return 0
	}

	if root.Val >= low {
		dfs(root.Left, low, high, result)
	}
	if root.Val >= low && root.Val <= high {
		*result += root.Val
	}
	if root.Val <= high {
		dfs(root.Right, low, high, result)
	}
	return 0
}

func main() {
	node := []int{10, 5, 15, 3, 7, 13, 18, 1, NULL, 6}
	root := Ints2TreeNode(node)
	low, high := 6, 10
	sumBST := rangeSumBST(root, low, high)
	fmt.Println(sumBST)
}
