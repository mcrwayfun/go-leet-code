package _49_intersetion

/**
给定两个数组，编写一个函数来计算它们的交集。
time complexity:O(n)
space complexity:O(n)
*/
func intersection(nums1 []int, nums2 []int) []int {
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
