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
// 1:递归求出所有子树元素和，放入map中，k存储元素和，v存储出现次数
// 2:循环判断map中出现次数最多的k
func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	m := make(map[int]int)
	dfs(root, m)

	var count int
	var ret []int
	for k, v := range m {
		if v == count {
			ret = append(ret, k)
		} else if v > count { // 发现出现次数更多的元素和，更新count以及ret
			count = v
			ret = []int{k}
		}
	}

	return ret
}

func dfs(root *TreeNode, m map[int]int) int {
	if root == nil {
		return 0
	}

	sum := root.Val + dfs(root.Left, m) + dfs(root.Right, m)
	if v, ok := m[sum]; ok {
		m[sum] = v + 1
	} else {
		m[sum] = 1
	}

	return sum
}
