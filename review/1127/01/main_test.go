package main

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	nums := []int{1, 2, 3}
	newNums := append(nums)
	fmt.Println(newNums)

	nums[0] = 5
	fmt.Println(nums)
	fmt.Println(newNums)
}
