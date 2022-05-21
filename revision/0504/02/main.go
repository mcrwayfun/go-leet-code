package main

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	var visitedMap = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		if _, ok := visitedMap[num]; ok {
			return []int{i, visitedMap[num]}
		}
		visitedMap[nums[i]] = i
	}
	return []int{-1, -1}
}
