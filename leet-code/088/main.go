package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 { // m==0, nums1位空，直接将nums2迁移到nums1
		copy(nums1, nums2)
		return
	}

	i, j, z := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; z-- { // 从后向前
		if nums1[i] >= nums2[j] {
			nums1[z] = nums1[i]
			i--
		} else {
			nums1[z] = nums2[j]
			j--
		}
	}

	// 仅需判断j>=0的情况
	for ; j >= 0; z-- {
		nums1[z] = nums2[j]
		j--
	}
}
