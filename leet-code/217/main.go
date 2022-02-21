package main

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	m := make(map[int]int, len(nums))
	for _, v := range nums{
		m[v]++
		if m[v] > 1 {
			return true
		}
	}
	return false
}
