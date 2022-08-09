package main

import "math/rand"

type Solution struct {
	Data     []int
	Original []int
}

func Constructor(nums []int) Solution {
	return Solution{
		Data:     nums,
		Original: append([]int{}, nums...),
	}
}

func (s *Solution) Reset() []int {
	copy(s.Data, s.Original)
	return s.Data
}

// Shuffle time complexity: O(n)
// space complexity: O(1)
func (s *Solution) Shuffle() []int {
	var n = len(s.Data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
	return s.Data
}
