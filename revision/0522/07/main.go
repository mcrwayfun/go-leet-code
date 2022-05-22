package main

// time complexity: O(n)
// space complexity: O(n)
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	var indexMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		if _, ok := indexMap[sub]; ok {
			return []int{i, indexMap[sub]}
		}
		indexMap[nums[i]] = i
	}
	return []int{-1, -1}
}
