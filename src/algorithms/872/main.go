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
		if i < n && ints[i] != 0 {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != 0 {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	var s1, s2 []int
	dfs(root1, &s1)
	dfs(root2, &s2)

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i]{
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*s = append(*s, root.Val)
	}

	for _, v := range []*TreeNode{root.Left, root.Right} {
		dfs(v, s)
	}

	return
}
