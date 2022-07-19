package main

import (
	"fmt"
	"sort"
)

/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func merge(intervals [][]int) [][]int {}
*/

/**
题目解析，假设现在有区间[1,8], [2,4], [9, 10], [10,16]
对于二维数组intervals，intervals[i][j] 分别表示区间的start和end。
1：先将区间按照起始的index来升序排序，即按照intervals[i][0]来排序，可以得到上述的例子
2：每次记录下res的长度，如果长度为0，可以直接将区间添加到结果集，对于第一个区间[1,8]，
即可以直接添加到结果集。
3：如果res[n-1].end < internal.start, 可以直接将区间加入到结果集。举个例子，现在res中有
区间[1,8]，如果区间[1,8].end = 8 < [9,10].start, 则两个区间不重合，可以直接将区间[9,10]
加入到结果集
4：否则两个区间有重叠的，需要进行区间的合并。合并区间是要取区间范围大的，即两个区间的end值做比较，
取一个大值。将当前结果集最后一个区间，即res[n-1].end = max(res[n-1].end, interval.end)
 */
// time complexity: O(n*lgn)
// space complexity: O(1)
func merge(intervals [][]int) [][]int {
	// 区间start按照升序的方式排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	for _, interval := range intervals {
		var n = len(res)
		if n == 0 || res != nil && res[n-1][1] < interval[0] {
			res = append(res, interval)
		} else { // 合并区间
			res[n-1][1] = max(res[n-1][1], interval[1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(merge([][]int{
		{9, 10},
		{2, 4},
		{1, 8},
		{10, 16},
	}))
}
