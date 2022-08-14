package main

import "sort"

// time complexity: O(n^2)
// space complexity: O(1)
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	sort.Ints(nums)
	var n = len(nums)
	var res [][]int
	for k := n - 1; k > 2; k-- {
		if nums[k] < 0 {
			break
		}

		var target = -nums[k]
		var i = 0
		var j = k - 1
		for i < j {
			if nums[i]+nums[j] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})

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

	return res
}
