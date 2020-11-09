package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nums []int

func kthSmallest(root *TreeNode, k int) int {
	// 二叉搜索树中序遍历是 升序的
	// 第k个最小元素 -> 升序中第k个元素
	nums = make([]int, 0)
	dfs(root)
	return nums[k-1]
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)
	nums = append(nums, root.Val)
	dfs(root.Right)
	return
}
