package main

import "math/rand"

type Solution struct {
	Nums     []int
	Original []int
}

func Constructor(nums []int) Solution {
	return Solution{
		Nums:     nums,
		Original: append([]int{}, nums...),
	}
}

func (s *Solution) Reset() []int {
	copy(s.Nums, s.Original)
	return s.Nums
}

func (s *Solution) Shuffle() []int {
	for i := len(s.Nums) - 1; i > 0; i-- {
		j := rand.Intn(i + 1) // [0, i+1)
		s.Nums[i], s.Nums[j] = s.Nums[j], s.Nums[i]
	}
	return s.Nums
}
