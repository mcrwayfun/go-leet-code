package main

import (
	. "../tool/tree"
	"fmt"
	"math"
)

type ResultType struct {
	Root2AnyPath int
	Any2AnyPath  int
}

// time complexity: O(n^2)
// space complexity: O(1)
func maxPathSum(root *TreeNode) int {
	return helper(root).Any2AnyPath
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{math.MinInt64, math.MinInt64}
	}

	left := helper(root.Left)
	right := helper(root.Right)

	// 根节点到任意节点的最大路径和
	root2AnyPath := max(0, max(left.Root2AnyPath, right.Root2AnyPath)) + root.Val

	// 任意节点到任意节点的最大路径和: 分为三个部分：根节点、左子树、右子树. 这里使用了两次分治法

	// 求左子树和右子树中any2any的最大值
	any2anyPath := max(left.Any2AnyPath, right.Any2AnyPath)
	// max(left.Root2AnyPath, 0) + max(right.Root2AnyPath, 0) + root.Val: 求根节点any2any的最大值
	any2anyPath = max(any2anyPath, max(left.Root2AnyPath, 0)+
		max(right.Root2AnyPath, 0)+root.Val)

	return ResultType{root2AnyPath, any2anyPath}
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
