package main

import "fmt"

// 给一个排序的数组和一个sum，求数组中的一对数，使它们的和 与 sum 相等，返回这对数的下标。
// arr=[1, 2, 3, 4, 6], target=6
func search(arr []int, targetSum int) []int {
	if len(arr) == 0 {
		return []int{-1, -1}
	}

	start := 0
	end := len(arr) - 1
	for start < end {
		sum := arr[start] + arr[end]
		if sum == targetSum {
			return []int{start, end}
		} else if sum > targetSum {
			end--
		} else {
			start++
		}
	}

	return []int{-1, -1}
}

func main() {
	arr := []int{1, 2, 3, 4, 6}
	targetSum := 6
	fmt.Println(search(arr, targetSum))
}
