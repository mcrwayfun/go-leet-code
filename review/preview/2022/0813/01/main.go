package main

import (
	"container/heap"
	"sort"
)

// time complexity: O(n^2)
// space complexity: O(n^2)
func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}

	var heights [][]int // 左边坐标的高度使用负数来表示
	for _, building := range buildings {
		heights = append(heights, []int{building[0], -building[2]})
		heights = append(heights, []int{building[1], building[2]})
	}

	// 对坐标集进行排序，先按照横坐标进行升序，横坐标相同按照纵坐标进行升序
	sort.Slice(heights, func(i, j int) bool {
		if heights[i][0] == heights[j][0] {
			return heights[i][1]-heights[j][1] < 0
		}
		return heights[i][0]-heights[j][0] < 0
	})

	// 使用最大堆来存储扫描过程的结果集
	// 每当最大值发生改变时，就产生了一个关键点
	var maxHeap = &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, 0)

	var preHeight int
	var res [][]int
	// 左边的坐标需要将纵坐标加入到结果集
	// 右边的坐标需要将纵坐标从结果集中移除
	for _, h := range heights {
		if h[1] < 0 {
			heap.Push(maxHeap, -h[1])
		} else {
			heap.Remove(maxHeap, maxHeap.FindIndex(h[1]))
		}

		// 操作结果集后，纵坐标最大值是否发生变化
		if preHeight != maxHeap.Peek().(int) {
			res = append(res, []int{h[0], maxHeap.Peek().(int)})
			preHeight = maxHeap.Peek().(int)
		}
	}

	return res
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

func (m MaxHeap) FindIndex(val int) int {
	if len(m) == 0 {
		return -1
	}

	for i, v := range m {
		if v == val {
			return i
		}
	}
	return -1
}

func (m MaxHeap) Peek() interface{} {
	return m[0]
}
