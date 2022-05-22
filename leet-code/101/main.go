package main

/**
题目：

给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/symmetric-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool
func isSymmetricTreeRecursive(root *TreeNode) bool
func isSymmetricTreeIterative(root *TreeNode) bool
*/

/**
题目解析：要求判断一颗二叉树是不是对称的，可以得到关键点为
假设左子树为s，右子树为t
1：s != null && t != null, 递归判断
	1-1：根节点上的值是否相等
	1-2：左节点的左子树和右节点的右子树是否相等
	1-3：左节点的右子树和右节点的左子树是否相等
2: 否则判断 s==null && t==null

Time Complexity: O(n)
Space Complexity: O(n) // 最坏情况只存在左右子树
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricHelper(s *TreeNode, t *TreeNode) bool {
	if s != nil && t != nil {
		return s.Val == t.Val &&
			isSymmetricHelper(s.Left, t.Right) && isSymmetricHelper(s.Right, t.Left)
	}
	return s == nil && t == nil
}

func isSymmetricTreeRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelper(root.Left, root.Right)
}

/**
将递归的方法转换为普通的循环方法，需要使用队列来实现
1：申明一个队列queue，假设左子树为s，右子树为t
2：将s和t入队列
3：只要队列不为空，则不断执行操作
	3-1：将队列顶两个元素pop出来，分别为s和t
	3-2：如果s == null && t == null continue
	3-3：如果s == null || t == null return false
	3-4：如果s.val != t.val return false
	3-5: queue.push(s.left), queue.push(t.right), queue.push(s.right), queue.push(t.left)
4: 结束可以返回true（没有找到不相等的情况）
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
