package _50_intersection_II

/**
给定两个数组，编写一个函数来计算它们的交集。
说明：
输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。我们可以不考虑输出结果的顺序。
给定两个数组，编写一个函数来计算它们的交集。

链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
time complexity:O(n)
space complexity:O(n)
*/
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	intersectMap := make(map[int]int)
	for _, v := range nums1 {
		intersectMap[v]++
	}

	ret := make([]int, 0)
	for _, v := range nums2 {
		if vv, ok := intersectMap[v]; ok && vv > 0 {
			ret = append(ret, v)
			intersectMap[v]--
		}
	}

	return ret
}
