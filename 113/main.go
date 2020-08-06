package main

import "fmt"

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
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var s []int
	dfs(root, 0, sum, &res, &s)
	return res
}

func dfs(root *TreeNode, tsum, sum int, res *[][]int, s *[]int) {
	if root == nil {
		return
	}

	tsum += root.Val
	*s = append(*s, root.Val)

	if root.Left == nil && root.Right == nil {
		if tsum == sum {
			*res = append(*res, append([]int{}, *s...))
		}
	}

	for _, v := range []*TreeNode{root.Left, root.Right} {
		dfs(v, tsum, sum, res, s)
	}

	// 若当前节点不符合，则将其从s中移除
	*s = (*s)[:len(*s)-1]
	return
}

func main() {
	node := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}
	root := int2TreeNode(node)
	sum := pathSum(root, 22)
	fmt.Println(sum)

	t := []int{1, 2, 3}
	_ = t
	fmt.Println(append([]int{}, t...))
}
