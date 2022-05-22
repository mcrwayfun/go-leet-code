package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
使用递归的方式来判断树是否对称
1：根节点为nil，肯定对称
2：只有根节点，肯定对称
3：如果左右子树非空，则递归判断左右子树是否对称
4：否则判断左右子树是否同时为空（一边非空则肯定不对称）

time complexity: O(n)
space complexity: O(n)
*/
func isSymmetricTreeRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetric(root.Left, root.Right)
}

func isSymmetric(s, t *TreeNode) bool {
	if s != nil && t != nil {
		return s.Val == t.Val && isSymmetric(s.Left, t.Right) && isSymmetric(s.Right, t.Left)
	}
	return s == nil && t == nil
}

/**
使用非递归的方式来判断树是否对称，需要借助辅助数据结构"队列"。与递归区别在于，
非递归是优先扫除非对称的情况
1：如果根节点为空，则肯定对称
2：初始化队列，将当前的根节点的左右子树入队
3：如果左右子树都为空，则继续循环
4：如果左右子树存在一个为空，则返回false
5：如果左右子树值不一致，则返回false
3：否则，左右子树均非空，则继续将它们入队。入队的顺序为
	s.Left, t.Right, s.Right, t.Left

time complexity: O(n)
space complexity: O(n)
 */
func isSymmetricTreeIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue = []*TreeNode{root.Left, root.Right}
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
		// s和t均为非空
		queue = append(queue, s.Left, t.Right, s.Right, t.Left)
	}

	return true
}
