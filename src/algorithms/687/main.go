package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func int2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != -1 {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != -1 {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	// 递归左右子树
	left := dfs(root.Left, res)
	right := dfs(root.Right, res)

	var leftLen,rightLen int
	if root.Left != nil && root.Left.Val == root.Val {
		leftLen = left + 1
	}

	if root.Right != nil && root.Right.Val == root.Val {
		rightLen = right + 1
	}

	*res = max(*res, leftLen + rightLen)
	return max(leftLen, rightLen)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}