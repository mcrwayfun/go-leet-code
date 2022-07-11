package main

import (
	"fmt"
	"sort"
)

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = []
输出：[]

示例 3：
输入：nums = [0]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func threeSum(nums []int) [][]int {}
*/

/**
1：先将数组排序
2：固定最后一个元素，对于[0,n-1]进行以下步骤：
3：如果 nums[i] + nums[j] == nums[k], ret.add。
	3-1：i移动到和nums[i]元素不相同的下标
	3-2：j移动到和nums[j]元素不相同的下标
4：nums[i] + nums[j] < nums[k], i++
5：nums[i] + nums[j] > nums[k], j--
6：k需要移动到nums[k]元素不相同给的下标

time complexity: O(n^2)
space complexity: O(1)
*/
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	sort.Ints(nums)
	var ret [][]int
	for k := len(nums) - 1; k >= 2; k-- {
		if nums[k] < 0 {
			break
		}

		var target = -nums[k]
		var i = 0
		var j = k - 1
		for i < j {
			if nums[i]+nums[j] == target {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})

				for i < j && nums[i+1] == nums[i] {
					i++
				}
				for i < j && nums[j-1] == nums[j] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				j--
			}

			for k >= 2 && nums[k-1] == nums[k] {
				k--
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
