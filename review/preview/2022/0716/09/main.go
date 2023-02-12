package main

import "sort"

// time complexity: O(nlgn)
// space complexity: O(1)
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var res [][]int
	sort.Ints(nums)
	for k := len(nums) - 1; k >= 2; k++ {
		if nums[k] < 0 {
			break
		}
		target := -nums[k]
		var i int
		var j = k - 1
		for i < j {
			if nums[i]+nums[j] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
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

	return res
}
