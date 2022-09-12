package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetric(root.Left, root.Right)
}

func isSymmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSymmetric(p.Left, q.Right) &&
		isSymmetric(p.Right, q.Left)
}

func isSymmetricIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack []*TreeNode
	stack = append(stack, root.Left, root.Right)

	for len(stack) > 0 {
		s := stack[0]
		stack = stack[1:]
		t := stack[0]
		stack = stack[1:]

		if s == nil && t == nil {
			continue
		}
		if s == nil || t == nil {
			return false
		}
		if s.Val != t.Val {
			return false
		}

		stack = append(stack, s.Left, t.Right, s.Right, t.Left)
	}

	return true
}
