package main

import "fmt"

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func findKthLargest(nums []int, k int) int {}
*/

/**
使用快速选择法来解答，根据题意，需要降序排序，求第k大的，等同于求数组中下标为k-1的元素，快速选择法的逻辑如下：
1：选择一个基值pivot，定义两个指针i和j，默认选择指针i对应的元素作为基准值
2：for nums[j] < pivot then j--
3：swap nums[i] and nums[j]
4：for nums[i] >= pivot then i++
5：swap nums[i] and nums[j]
最后得到一个降序的数组

当原数组是升序的，最坏的时间复杂度为 O(n^2)，可以使用洗牌算法打乱数组，使得时间复杂度降到O(n)
平均时间复杂度为 O(n)
*/
func findKthLargest(nums []int, k int) int {
	var low = 0
	var high = len(nums) - 1

	for low <= high {
		p := partition(nums, low, high)
		if p == k-1 {
			return nums[p]
		} else if p < k-1 {
			low = p + 1
		} else {
			high = p - 1
		}
	}

	return -1
}

func partition(nums []int, low, high int) int {
	var pivot = nums[low]
	var i = low
	var j = high

	for i < j {
		for ; i < j && nums[j] < pivot; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]
		for ; i < j && nums[i] >= pivot; i++ {
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i // i == j
}

func main() {
	num := []int{3, 2, 1, 5, 6, 4}
	ans := findKthLargest(num, 2)
	fmt.Println(ans)
}
