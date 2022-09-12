package main

import "sort"

// time complexity: O(n^2)
// space complexity: O(1)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for k := len(nums) - 1; k >= 2; k-- {
		if nums[k] < 0 {// [0,0,0] 也是答解
			break
		}

		var i = 0
		var j = k - 1
		var target = -nums[k]
		for i < j {
			if nums[i]+nums[j] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				for i < j && nums[i] == nums[i+1] {
					i++// skip the same ele
				}
				for i < j && nums[j] == nums[j-1] {
					j-- // skip the same ele
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				j--
			}
		}

		for k > 2 && nums[k-1] == nums[k] {
			k--
		}
	}

	return res
}
