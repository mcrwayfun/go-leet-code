package main

import (
	"container/heap"
	"sort"
)

// time complexity: O(n^2)
// space complexity: O(n^2)
func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 || len(buildings[0]) == 0 {
		return [][]int{}
	}

	var height [][]int
	for _, b := range buildings {
		height = append(height, []int{b[0], -b[2]})
		height = append(height, []int{b[1], b[2]})
	}

	// 按照横坐标升序排序，横坐标相同按照纵坐标升序
	sort.Slice(height, func(i, j int) bool {
		if height[i][0] == height[j][0] {
			return height[i][2]-height[j][2] < 0
		}
		return height[i][0]-height[j][0] < 0
	})

	// 使用最大堆来存储坐标的结果集
	var maxHeap = &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, 0)

	var result [][]int
	var preHeight int // 之前的高度
	for _, h := range height {
		if h[1] < 0 {
			heap.Push(maxHeap, -h[1])
		} else {
			heap.Remove(maxHeap, maxHeap.FindIndex(h[1]))
		}

		if preHeight != maxHeap.Peek().(int) {// 高度发生了变化，产生了关键点
			result = append(result, []int{h[0], maxHeap.Peek().(int)})
			preHeight = maxHeap.Peek().(int)
		}
	}
	return result
}

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m *MaxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return res
}

func (m MaxHeap) Peek() interface{} {
	return m[0]
}

func (m MaxHeap) FindIndex(val int) int {
	for i, v := range m {
		if v == val {
			return i
		}
	}
	return -1
}
