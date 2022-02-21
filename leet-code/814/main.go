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
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 后序遍历的方式剪掉所有值为 0 的叶子节点
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil // 剪支
	}

	return root
}