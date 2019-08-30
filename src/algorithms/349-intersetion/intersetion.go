package _49_intersetion

/**
给定两个数组，编写一个函数来计算它们的交集。
说明:
输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。

链接：https://leetcode-cn.com/problems/intersection-of-two-arrays

time complexity:O(n)
space complexity:O(n)
*/
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	// 使用map存储nums1
	temp := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		temp[v] = 1
	}

	temp2 := make(map[int]int, len(temp))
	for _, v := range nums2 {
		// map中存在该元素
		if _, ok := temp[v]; ok {
			temp2[v] = v
		}
	}

	var ret []int
	for _, v := range temp2 {
		ret = append(ret, v)
	}
	return ret
}

func intersection2(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	// 使用map存储nums1
	temp := make(map[int]bool, len(nums1))
	for _, v := range nums1 {
		temp[v] = false
	}

	var ret []int
	for _, v := range nums2 {
		// map中存在该元素
		if vv, ok := temp[v]; ok && !vv {
			ret = append(ret, v)
			temp[v] = true
		}
	}

	return ret
}