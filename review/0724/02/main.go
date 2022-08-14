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
	for _, building := range buildings {
		height = append(height, []int{building[0], -building[2]})
		height = append(height, []int{building[1], building[2]})
	}

	// 横坐标相同，按照纵坐标排序，否则按照纵坐标排序（升序）
	sort.Slice(height, func(i, j int) bool {
		if height[i][0] == height[j][0] {
			return height[i][1]-height[j][1] < 0
		}
		return height[i][0]-height[j][0] < 0
	})

	var result [][]int
	var preHeight int // 上一次的最大高度
	var maxHeap = &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, 0)

	for _, h := range height {
		if h[1] < 0 { // 矩形的左边
			heap.Push(maxHeap, -h[1])
		} else { // 矩形的右边，移除高度
			heap.Remove(maxHeap, maxHeap.FindIndex(h[1]))
		}

		if maxHeap.Peek() != preHeight { // 高度发生变化，产生了关键点
			result = append(result, []int{h[0], maxHeap.Peek().(int)})
			preHeight = maxHeap.Peek().(int)
		}
	}

	// 使用最大堆来存储结果集
	return result
}

type MaxHeap []int

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Peek() interface{} {
	return h[0]
}

func (h MaxHeap) FindIndex(val int) int {
	for i, v := range h {
		if v == val {
			return i
		}
	}
	return -1
}
