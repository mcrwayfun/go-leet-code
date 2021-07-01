package main

import . "../tool/tree"

/**
1：只要碰到p或q，就直接返回
2：找到了LCA就返回
	2-1: left和right都存在，说明祖先是root
	2-2：left不为空则返回left
	2-3：right不为空则返回right
3：如果都没有，就返回NULL
*/
// time complexity: O(n)
// space complexity: O(1)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil
}
