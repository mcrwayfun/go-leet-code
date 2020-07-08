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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}

	for _, v := range []*TreeNode{root.Left, root.Right} {
		sum += sumOfLeftLeaves(v)
	}

	return sum
}