package main

import "fmt"

// 给定一个正数数组和正数S，求最小的连续子数组的长度，其总和大于或等于S，
// 如果不存这样的子数组则返回0
// for windowEnd-windowStart+1>=S 则滑动窗口
func findMinSubArray(s int, arr []int) int {
	if len(arr) == 0 || s == 0 {
		return 0
	}

	var windowStart int
	var windowSum int
	var minLength = len(arr) + 1

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]

		// 滑动窗口后可能会存在 windowSum >= s
		// 所以需要继续滑动
		for windowSum >= s {
			if windowEnd-windowStart+1 < minLength {
				minLength = windowEnd - windowStart + 1
			}
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	if minLength == len(arr)+1 {
		return 0
	}

	return minLength
}

func main() {
	arr := []int{2, 1, 5, 2, 3, 2}
	s := 7
	fmt.Println(findMinSubArray(s, arr))
}
