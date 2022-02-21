package main

import (
	. "../tool/tree"
	"fmt"
)

// 中序遍历：左中右
// time complexity: O(n)
// space complexity: O(1)
func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	// divide
	leftRet := inorderTraversal(root.Left)
	rightRet := inorderTraversal(root.Right)

	// conquer
	result = append(result, leftRet...)
	result = append(result, root.Val)
	result = append(result, rightRet...)
	return result
}

// 非递归中序遍历
// time complexity: O(n)
// space complexity: O(n)
func inorderTraversalNonRecursion(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}

	return result
}

func main() {
	nums := []int{1, NULL, 2, 3}
	root := Ints2TreeNode(nums)

	ret := inorderTraversal(root)
	fmt.Println("中序递归的结果为:", ret)

	ret = inorderTraversalNonRecursion(root)
	fmt.Println("中序非递归的结果为:", ret)
}
