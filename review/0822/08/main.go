package main

// time complexity: O(m*n)
// space complexity: O(1)
func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return 0
	}

	var m = len(haystack)
	var n = len(needle)
	for i := 0; i <= m-n; i++ {
		var k = i
		var z = 0
		for ; z < n && haystack[k] == needle[z]; {
			k++
			z++
		}
		if z == n {
			return i
		}
	}
	return -1
}
