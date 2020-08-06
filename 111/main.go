package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}

	if root.Left == nil && root.Right == nil{// 叶子节点
		return 1
	}
	leftMinDepth := minDepth(root.Left)
	rightMinDepth := minDepth(root.Right)

	// 左节点或者右节点为空的情况，只返回不为空的节点的深度
	if root.Left == nil || root.Right == nil{
		return leftMinDepth + rightMinDepth + 1
	}

	if leftMinDepth > rightMinDepth{
		return rightMinDepth + 1
	}

	return leftMinDepth + 1
}