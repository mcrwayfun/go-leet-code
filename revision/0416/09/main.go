package main

import (
	"fmt"
	"sort"
)

// 有一个未排序的数组和target sum，求所有满足 arr[i] + arr[j] + arr[k] < target 的下标，
// 写一个函数返回满足条件的 三数之和 的数量。
// Input: [-1, 0, 2, 3], target=3
// Output: 2
// Input: [-1, 1, 2, 3, 4], target=5
// Output: 4
//
func searchTriplets(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)

	var count int
	for i := 0; i < len(arr); i++ {
		j := i + 1
		k := len(arr) - 1

		for j < k {
			sum := arr[i] + arr[j] + arr[k]
			if sum < target {
				count += k - j
				j++
			} else {
				k--
			}
		}
	}

	return count
}

func main() {
	arr := []int{-1, 4, 2, 1, 3}
	target := 5
	fmt.Println(searchTriplets(arr, target))
}
