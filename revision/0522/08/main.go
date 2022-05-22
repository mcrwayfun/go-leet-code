package main

/**
判断两颗二叉树是否完全相等，使用递归的方法来判断
1：如果p和q均非空，则递归判断左右子树是否相等
2：否则判断p和q是否均为空

time complexity: O(n)
space complexity: O(n)
*/
func isSameTreeRecursive(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTreeRecursive(p.Left, q.Left) &&
			isSameTreeRecursive(p.Right, q.Right)
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

		queue = append(queue, s.Left, t.Left, s.Right, t.Right)
	}
	return true
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
