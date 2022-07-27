package main

// time complexity: O(n)
// space complexity: O(n)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	return hasPathSum(root.Left, targetSum - root.Val) ||
		hasPathSum(root.Right, targetSum - root.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
