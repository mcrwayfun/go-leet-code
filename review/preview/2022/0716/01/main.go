package main

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	var visitedMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		if _, ok := visitedMap[sub]; ok {
			return []int{i, visitedMap[sub]}
		}

		visitedMap[nums[i]] = i
	}
	return []int{-1, -1}
}
