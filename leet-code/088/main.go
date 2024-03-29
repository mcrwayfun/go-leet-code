package main

/**
题目：
给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。
示例 3：

输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：
func merge(nums1 []int, m int, nums2 []int, n int) {}
*/

/**
解题思路：由于nums1的长度足够，可以将nums2中的元素全部搬到nums1中。
如果下标从前向后遍历，最坏情况下nums1中的元素需要向后移动，所以我们
从nums1和nums2的末尾开始遍历，这样可以减少搬迁。
1：申明下标i=m-1,j=n-1,k=m+n-1 (k是搬迁存储元素的下标）
2：如果nums1[i] > nums2[j], nums1[k]=nums1[i], i--,k--
3：否则nums1[k]=nums2[j]，j--,k--
4：遍历结束后，如果j>=0，将剩余的元素全部搬迁到nums1即可

time complexity: O(m+n)
space complexity: O(1)
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
