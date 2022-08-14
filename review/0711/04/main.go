package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	sort.Ints(nums)
	var ret [][]int
	for k := len(nums) - 1; k >= 2; k-- {
		if nums[k] < 0 {
			break
		}

		var i = 0
		var j = k - 1
		var target = -nums[k]

		for i < j {
			if nums[i]+nums[j] == target {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})

				for i < j && nums[i+1] == nums[i] {
					i++
				}
				for i < j && nums[j-1] == nums[j] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				j--
			}
		}

		for k >= 2 && nums[k-1] == nums[k] {
			k--
		}
	}

	return ret
}
