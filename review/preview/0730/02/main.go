package main

// time complexity: O(n)
// space complexity: O(n)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue = []*TreeNode{root}
	var s *TreeNode
	var ret [][]int
	for len(queue) > 0 {
		var level []int
		size := len(queue)
		for i := 0; i < size; i++ {
			s, queue = queue[0], queue[1:]
			level = append(level, s.Val)

			if s.Left != nil {
				queue = append(queue, s.Left)
			}
			if s.Right != nil {
				queue = append(queue, s.Right)
			}
		}
		ret = append(ret, level)
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
