package main

/// Time Complexity: O( len(nums1) + len(nums2) )
/// Space Complexity: O( len(nums1) )
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)

	for _, v := range nums1 {
		m[v] = false
	}

	m2 := make(map[int]int)
	for _, v := range nums2 {
		if _,ok := m[v];ok {
			m2[v] = v
		}
	}

	var ret []int
	for _, v := range m2{
		ret = append(ret, v)
	}

	return ret
}
