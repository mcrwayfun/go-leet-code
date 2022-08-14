package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time complexity: O(n)
// space complexity: O(n)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	}
	return p == nil && q == nil
}

// time complexity: O(n)
// space complexity: O(n)
func isSameTreeIteration(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	var queue = []*TreeNode{p, q}
	var s *TreeNode
	var t *TreeNode
	for len(queue) > 0 {
		s, queue = queue[0], queue[1:]
		t, queue = queue[0], queue[1:]

		if s == nil && t == nil {
			continue
		}
		if s == nil || t == nil {
			return false
		}
		if s.Val != t.Val {
			return false
		}

		// s != nil && t != nil
		queue = append(queue, s.Left, t.Left, s.Right, t.Right)
	}

	return true
}
