package main

// time complexity: O(n)
// space complexity: O(1)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var res int
	var b = x
	for b != 0 {
		tmp := b%10
		res = res * 10 + tmp
		b = b/10
	}
	return res == x
}
