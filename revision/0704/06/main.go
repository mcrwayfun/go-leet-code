package main

func hammingWeight(num uint32) int {
	var n = num
	var count int
	for n != 0 {
		count++
		n = n & (n-1)
	}
	return count
}
