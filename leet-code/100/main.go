package main

/**
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1：
输入：p = [1,2,3], q = [1,2,3]
输出：true

示例 2：
输入：p = [1,2], q = [1,null,2]
输出：false

示例 3：
输入：p = [1,2,1], q = [1,1,2]
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/same-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func isSameTree(p *TreeNode, q *TreeNode) bool
*/

/**
解题思路：使用递归的方式，判断
1：左右子树同时为nil，相等
2：左右子树有一个为nil，不相等
3：根节点值不相等
4：递归判断根节点左右子树是否相等

time complexity: O(n)
space complexity: O(n) // 最坏的情况
*/
func isSameTreeRecursive(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTreeRecursive(p.Left, q.Left) &&
		isSameTreeRecursive(p.Right, q.Right)
}

/**
解题思路：使用循环的方式来判断，需要申请额外的数据结构队列
1：先写结束条件，如果左右子树同时为空，直接返回true
2：申请额外的数据结构队列，将当前左右子树push进队列
3：如果队列不为空，则执行以下的步骤
	3-1：依次pop队列顶元素s和t
	3-2：if s == nil && t == nil continue // 当前相等，循环判断
	3-3：if s == nil || t == nil return false
	3-4：if s.val != t.val return false
	3-5：queue.push(s.left), queue.push(t.left), queue.push(s.right), queue.push(t.right)

time complexity: O(n)
space complexity: O(n)
*/
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
