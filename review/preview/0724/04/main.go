package main

// time complexity: O(n)
// space complexity: O(n)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left != nil && root.Right != nil {
		return root.Val == targetSum
	}

	return hasPathSum(root.Left, root.Val-targetSum) ||
		hasPathSum(root.Right, root.Val-targetSum)
}

// time complexity: O(n)
// space complexity: O(n)
func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var stack = []*TreeNode{root}
	var sumStack = []int{targetSum}
	var s *TreeNode
	var n int
	for len(stack) > 0 {
		s, stack = stack[len(stack)-1], stack[:len(stack)-1]
		n, sumStack = sumStack[len(sumStack)-1], sumStack[:len(sumStack)-1]

		if s.Left == nil && s.Right == nil && s.Val == n {
			return true
		}

		if s.Left != nil {
			stack = append(stack, s.Left)
			sumStack = append(sumStack, n - s.Val)
		}

		if s.Right != nil {
			stack = append(stack, s.Right)
			sumStack = append(sumStack, n - s.Val)
		}
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
