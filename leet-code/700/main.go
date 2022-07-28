package main

/**
给定二叉搜索树（BST）的根节点root和一个整数值val。
你需要在 BST 中找到节点值等于val的节点。 返回以该节点为根的子树。 如果节点不存在，则返回null。

示例 1:
输入：root = [4,2,7,1,3], val = 2
输出：[2,1,3]

示例 2:
输入：root = [4,2,7,1,3], val = 5
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/search-in-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func searchBST(root *TreeNode, val int) *TreeNode {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

// time complexity: O(lgn)
// space complexity: O(h)
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

// time complexity: O(lgn)
// space complexity: O(1)
func searchBSTIterative(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
