package main

// time complexity: O(n)
// space complexity: O(1)
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{-1, -1}
	}

	var left = 0
	var right = len(numbers) - 1
	for left <= right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
