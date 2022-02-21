package main

import (
	. "../tool/tree"
	"fmt"
	"math"
)

type ResultType struct {
	Root2Any int
	Any2Any int
}

// time complexity: O(n^2)
// space complexity: O(1)
func maxPathSum(root *TreeNode) int {
	return helper(root).Any2Any
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{math.MinInt64, math.MinInt64}
	}

	left := helper(root.Left)
	right := helper(root.Right)

	root2Any := max(0, max(left.Root2Any, right.Root2Any)) + root.Val // 从根节点->任意节点的最大路径和（仅有一边）

	// 从任意节点->任意节点的最大路径和
	any2Any := max(left.Any2Any, right.Any2Any) // 左右子树的
	// 从根节点->任意节点的最大路径和（包含根节点在内，两边都要算上）
	any2Any = max(0, max(0, left.Root2Any) + max(0, right.Root2Any) + root.Val)
	return ResultType{root2Any, any2Any}
}

var maxPath = -1001

func maxPathSumV2(root *TreeNode) int {
	maxPath = 0
	helperV2(root)
	return maxPath
}

func helperV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := helperV2(root.Left)
	if left <= 0 {
		left = 0
	}

	right := helperV2(root.Right)
	if right <= 0 {
		right = 0
	}

	maxPath = max(maxPath, left+right+root.Val)
	return max(root.Val+left, root.Val+right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 从根节点开始到任意节点，求最大路径和
func maxPathSumII(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftPath := maxPathSumII(root.Left)
	rightPath := maxPathSumII(root.Right)

	// max(leftPath, rightPath) 当前树的左右节点中选一个最大值
	// max(0, max(leftPath, rightPath)) 选出的最大值，若小于0则不要
	// max(0, max(leftPath, rightPath)) + root.Val 选出最大值，至少要包含根节点
	return max(0, max(leftPath, rightPath)) + root.Val
}

func main() {
	root := Ints2TreeNode([]int{-10, 9, 20, NULL, NULL, 15, 7})
	sum := maxPathSum(root)
	fmt.Println(sum)
}
