package main

/**
Given the root of a binary tree, invert the tree, and return its root.

Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/invert-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func invertTree(root *TreeNode) *TreeNode {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

// time complexity: O(n)
// space complexity: O(n)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// time complexity: O(n)
// space complexity: O(1)
func invertTreeIterator(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var queue = []*TreeNode{root}
	var s *TreeNode
	for len(queue) > 0 {
		s, queue = queue[0], queue[1:]

		s.Left, s.Right = s.Right, s.Left
		if s.Left != nil {
			queue = append(queue, s.Left)
		}
		if s.Right != nil {
			queue = append(queue, s.Right)
		}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
