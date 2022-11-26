package main

// time complexity: O(m+n)
// space complexity: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	var k = m + n - 1
	var i = m - 1
	var j = n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

	return
}
