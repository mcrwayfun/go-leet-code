package main

import (
	"fmt"
	"sort"
)

// 给定一个未排序的数组，找到所有唯一的 三个数加起来等于0的。
// 先排序，固定一个位置，满足 arr[i] = arr[j]+arr[k]
// arr=[-3, 0, 1, 2, -1, 1, -2]
// sort arr=[-3,-2,-1,0,1,1,2]
func searchTriplets(arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}

	sort.Ints(arr)

	var triplets [][]int
	for i := 0; i < len(arr); i++ {
		if i > 1 && arr[i] == arr[i-1] {
			continue
		}
		search(arr, arr[i], i+1, &triplets)
	}

	return triplets
}

func search(arr []int, target int, start int, triplets *[][]int) {
	var end = len(arr) - 1

	for start < end {
		sum := target + arr[start] + arr[end]
		if sum == 0 {
			*triplets = append(*triplets, []int{target, arr[start], arr[end]})

			start++
			end--
			for start < end && arr[start] == arr[start-1] {
				start++
			}
			for start < end && arr[end] == arr[end+1] {
				end--
			}
		} else if sum > 0 {
			end--
		} else {
			start++
		}
	}
}

func main() {
	arr := []int{-3, 0, 1, 2, -1, 1, -2}
	fmt.Println(searchTriplets(arr))
}
