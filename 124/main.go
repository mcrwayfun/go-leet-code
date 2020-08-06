package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func int2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != -1 {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != -1 {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res :=-1 << 31
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return -1 << 31
	}

	left := dfs(root.Left, res)
	right := dfs(root.Right, res)

	// 获取PathSum
	*res = max(*res, root.Val+max(0, left)+max(0, right))
	// 计算当前路径的最大值
	return root.Val + max(0, max(left, right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := int2TreeNode([]int{1, 2, 3})
	sum := maxPathSum(root)
	fmt.Println(sum)
}
