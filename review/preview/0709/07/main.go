package main

func hammingWeight(num uint32) int {
	var count int
	var n = int(num)
	for n != 0 {
		count++
		n = n & (n-1)
	}
	return count
}

func hammingWeightV2(num uint32) int {
	var count int
	for num != 0 {
		if num & 1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
