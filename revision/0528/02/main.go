package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
判断一颗二叉树是否对称
1：根节点为空则必定轴对称
2：根节点的左右子树都不为空，递归判断左右子树是否对称
3：根节点的左右子树是否同时为空，否则不对称

time complexity: O(n)
space complexity: O(n)
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricTreeRecursive(root.Left, root.Right)
}

func isSymmetricTreeRecursive(s *TreeNode, t *TreeNode) bool {
	if s != nil && t != nil {
		return s.Val == t.Val && isSymmetricTreeRecursive(s.Left, t.Right) &&
			isSymmetricTreeRecursive(s.Right, t.Left)
	}
	return s == nil && t == nil
}

/**
使用非递归的方式来判断：定义一个队列，入队的顺序为s.left, t.right, s.right, t.left

time complexity: O(n)
space complexity: O(n)
 */
func isSymmetricTreeIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue = []*TreeNode{root}
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

		queue = append(queue, s.Left, t.Right, s.Right, t.Left)
	}

	return true
}