package main

import "math/rand"

type Solution struct {
	Data   []int
	Source []int
}

func Constructor(nums []int) Solution {
	return Solution{
		Data: nums,
		// append(nums) 底层指向的还是同一个ptr
		// append([]int{}, nums...) 会生成一个新的切片ptr
		Source: append([]int{}, nums...),
	}
}

func (s *Solution) Reset() []int {
	copy(s.Data, s.Source)
	return s.Data
}

// Shuffle 洗牌算法
// [0, i] 保持顺序, (i, n-1] 无序
func (s *Solution) Shuffle() []int {
	for i := len(s.Data) - 1; i >= 0; i-- {
		var j = rand.Intn(i + 1)
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
	return s.Data
}
