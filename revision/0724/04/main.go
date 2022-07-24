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

	return hasPathSum(root.Left, root.Val - targetSum) ||
		hasPathSum(root.Right, root.Val - targetSum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
