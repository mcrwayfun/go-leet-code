package main

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{-1, -1}
	}

	var start int
	var end = len(numbers) - 1

	for start < end {
		num := numbers[start] + numbers[end]
		if num < target {
			start++
		} else if num > target {
			end--
		} else {
			return []int{start + 1, end + 1}
		}
	}
	return []int{-1, -1}
}
