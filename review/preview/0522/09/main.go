package main

/**
由于是已经按照递增顺序排好序的数组，可以使用两个指针来遍历数组，
以此来达到时间复杂度为O(n)的算法
*/
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{-1, -1}
	}

	var start, end = 0, len(numbers) - 1
	for start < end {
		num := numbers[start] + numbers[end]
		if num == target {
			return []int{start + 1, end + 1}
		} else if num < target {
			start++
		} else {
			end--
		}
	}
	return []int{-1, -1}
}
