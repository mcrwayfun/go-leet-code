package main

/**
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depthis the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
Input: root = [1,null,2]
Output: 2

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func maxDepth(root *TreeNode) int{}

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
*/

// time complexity: O(n)
// space complexity: O(n)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// time complexity: O(n)
// space complexity: O(n)
func maxDepthWithIteration(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue = []*TreeNode{root}
	var s *TreeNode
	var depth int

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			s, queue = queue[0], queue[1:]
			if s.Left != nil {
				queue = append(queue, s.Left)
			}
			if s.Right != nil {
				queue = append(queue, s.Right)
			}
		}
		depth++
	}

	return depth
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
