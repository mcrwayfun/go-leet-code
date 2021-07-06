package main

import (
	. "../tool/tree"
	"fmt"
)

// time complexity: O(n)
// space complexity: O(n)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var result []int
	dfs(root, &result)

	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, result)
	*result = append(*result, root.Val)
	dfs(root.Right, result)
	return
}

func main() {
	nums := []int{2, 1, 3}
	root := Ints2TreeNode(nums)
	validBST := isValidBST(root)
	fmt.Println(validBST)

	nums = []int{5, 1, 4, NULL, NULL, 3, 6}
	root = Ints2TreeNode(nums)
	validBST = isValidBST(root)
	fmt.Println(validBST)

	nums = []int{2, 2, 2}
	root = Ints2TreeNode(nums)
	validBST = isValidBST(root)
	fmt.Println(validBST)
}
