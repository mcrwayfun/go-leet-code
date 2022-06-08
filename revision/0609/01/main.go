package main

// time complexity: O(n)
// space complexity: O(n)
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}
	var s *TreeNode
	for len(queue) > 0 {
		size := len(queue)
		var level []int
		for i:=0; i<size; i++ {
			s, queue = queue[0], queue[1:]
			if s.Left != nil {
				queue = append(queue, s.Left)
			}
			if s.Right != nil {
				queue = append(queue, s.Right)
			}
			level = append(level, s.Val)
		}
		res = append(res, level)
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}