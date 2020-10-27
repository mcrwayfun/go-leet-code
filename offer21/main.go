package main

/*
思路1：使用新的数组来存储奇数，后面append原来的数组上
 */
func exchange(nums []int) []int {
	var oddNums []int
	var evenNums []int
	for _, v := range nums {
		if v % 2 == 0 {
			evenNums = append(evenNums, v)
		}else {
			oddNums = append(oddNums, v)
		}
	}

	oddNums = append(oddNums, evenNums...)
	return oddNums
}

func exchange2(nums []int) []int{
	start := 0
	end := len(nums) - 1
	for start < end && start < len(nums)-1 && end > 0 {
		if nums[start]%2 == 0 && nums[end]%2 != 0 {// start是偶数，end是奇数
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
		if nums[start]%2 == 0 && nums[end]%2 == 0 {// start是偶数，end也是偶数
			end--
		}
		if nums[start]%2 != 0 && nums[end]%2 != 0 {// start奇数，end也是奇数
			start++
		}
		if nums[start]%2 != 0 && nums[end]%2 == 0 {// start奇数，end是偶数
			start++
			end--
		}
	}
	return nums
}