package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**
城市的 天际线 是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，
请返回 由这些建筑物形成的 天际线 。

每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [lefti, righti, heighti] 表示：

lefti 是第 i 座建筑物左边缘的 x 坐标。
righti 是第 i 座建筑物右边缘的 x 坐标。
heighti 是第 i 座建筑物的高度。
你可以假设所有的建筑都是完美的长方形，在高度为 0 的绝对平坦的表面上。

天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序 。关键点是水平线段的左端点。列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

注意：输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...]
是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/the-skyline-problem
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func getSkyline(buildings [][]int) [][]int {}
*/

// time complexity: O(n^2)
// space complexity: O(n)
func getSkyline(buildings [][]int) [][]int {
	var result [][]int
	var height [][]int

	for _, b := range buildings {
		// 左边坐标和高度负值
		height = append(height, []int{b[0], -b[2]})
		// 右边坐标和高度正值
		height = append(height, []int{b[1], b[2]})
	}
	sort.Slice(height, func(i, j int) bool {
		if height[i][0] == height[j][0] {
			// 横坐标相等，按照纵坐标的高度值排序
			return height[i][1]-height[j][1] < 0
		}
		// 否则按照横坐标排序
		return height[i][0]-height[j][0] < 0
	})

	// 声明最大堆，用于保存高度值
	var maxHeap = &MaxHeap{}
	heap.Init(maxHeap)
	// 加入初始的高度0
	heap.Push(maxHeap, 0)

	// 上一次的最大高度
	var preHeight int
	// 遍历高度数组
	for _, h := range height {
		if h[1] < 0 { // 矩形的左边，将高度值加入堆中
			heap.Push(maxHeap, -h[1])
		} else { // 矩形的右边，移除高度值
			heap.Remove(maxHeap, maxHeap.FindIndex(h[1]))
		}

		// 最大高度发生变化，出现关键点
		if maxHeap.Peek().(int) != preHeight {
			result = append(result, []int{h[0], maxHeap.Peek().(int)})
			preHeight = maxHeap.Peek().(int)
		}
	}

	return result
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h[i] > h[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出最后一个元素
func (h *MaxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *MaxHeap) Peek() interface{} {
	res := (*h)[0]
	return res
}

func (h *MaxHeap) FindIndex(val int) int {
	for i, v := range *h {
		if v == val {
			return i
		}
	}
	return -1
}

func main() {
	/**
	1,-2
	3,2
	2,-3
	4,3
	5,-2
	9,2
	6,-4
	8,4

	排序：
	1，-2
	2，-3
	3，2
	4，3
	5，-2
	6，-4
	8，4
	9，2
	*/
	var buildings = [][]int{
		{1, 3, 2},
		{2, 4, 3},
		{5, 9, 2},
		{6, 8, 4},
	}
	fmt.Println(getSkyline(buildings))

	buildings = [][]int{
		{2, 4, 2},
		{2, 4, 3},
		{5, 7, 2},
		{7, 9, 4},
	}
	fmt.Println(getSkyline(buildings))
}
