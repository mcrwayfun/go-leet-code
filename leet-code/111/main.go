package main

/**
题目：
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：2

示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func minDepth(root *TreeNode) int
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
解法1：使用递归的方式来求树的最小深度
1：如果根节点为空，则直接返回0
2：如果左右节点均为空，则直接返回1（只有根节点）。
否则说明左右子树中至少有一棵非空
3：如果左子树为空，则递归去求右子树的最小深度+1
4：如果右子树为空，则递归去求左子树的最小深度+1
5：否则左右子树均非空，则递归求左右子树最小深度+1
+1表示当前层深度

time complexity: O(n)
space complexity: O(n)
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1 // 根节点
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

/**
解法2：使用非递归的方式来求最小深度
1：定义一个队列，将树的根节点入队
2：队列不为空则执行下列步骤
	2-1：当前需要遍历的子节点 size=queue.size
	2-2: 遍历子节点，如果子节点的左右子树均为空，则直接返回depth
	2-3：如果左子树非空，则左子树入队；右子树非空，右子树入队
	2-4：结束当前size的循环，depth++

time complexity: O(n)
space complexity: O(n)
*/
func minDepthWithIteration(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue = []*TreeNode{root}
	var node *TreeNode
	var depth = 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node, queue = queue[0], queue[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
