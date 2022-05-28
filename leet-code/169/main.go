package main

/**
给定一个大小为 n 的数组nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例1：
输入：nums = [3,2,3]
输出：3

示例2：
输入：nums = [2,2,1,1,1,2,2]
输出：2

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/majority-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func majorityElement(nums []int) int
*/

/**
假设数组中第一个元素是出现次数过半的元素num，用count来表示当前元素出现的次数
1：如果count==0, 则num=当前元素，并且count设置为1
2：如果 nums[i]==num, 则count++, 否则count--

time complexity: O(n)
space complexity: O(1)
*/
func majorityElement(nums []int) int {
	var num, count = nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			num = nums[i]
			count = 1
		} else if nums[i] == num {
			count++
		} else {
			count--
		}
	}
	return num
}
