package main

import "math/rand"

type Solution struct {
	Data   []int
	Source []int
}

func Constructor(nums []int) Solution {
	return Solution{
		Data:   nums,
		Source: append(nums),
	}
}

func (s *Solution) Reset() []int {
	copy(s.Data, s.Source)
	return s.Data
}

// Shuffle 洗牌算法本质上就是利用 rand 的方式来进行随机交换
// 将数组切分成两部分, [0, i] 是顺序的, (i, n-1] 是乱序的
func (s *Solution) Shuffle() []int {
	var n = len(s.Data)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1) // [0,i+1)
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
	return s.Data
}
