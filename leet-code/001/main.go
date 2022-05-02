package main

import "fmt"

/**
题目：

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func twoSum(nums []int, target int) []int
*/

/**
解题思路：从数组中找到两个数a+b=target，并返回这a和b的下标

解法1：使用暴力破解法
1：使用双层循环找到 a+b=target，找不到则返回-1,-1
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
func twoSumWithBrute(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

/**
解法2：
记录已经访问过的元素，这样可以不再需要遍历已经访问过的
1：使用辅助的数据结构map来存储已经访问过的元素，key存储访问过的元素，value存储它的下标
2：遍历数组，如果target-nums[i]存在于map中，说明访问过的元素中存在nums[i]+nums[j]==target，
可以返回下标i和map.get(target-nums[i])
3：否则 key设置为nums[i]，value为nums[i]的下标存储到map中
4：遍历结束仍然不满足则返回-1,-1
Time Complexity: O(n)
Space Complexity: O(n)
*/
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	visitedMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		if _, ok := visitedMap[num]; ok {
			return []int{visitedMap[num], i}
		}
		visitedMap[nums[i]] = i
	}

	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumWithBrute(nums, target))
	fmt.Println(twoSum(nums, target))
}
