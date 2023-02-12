package main

// time complexity: O(n^2)
// space complexity: O(1)
func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return 0
	}

	var m = len(haystack)
	var n = len(needle)
	for i := 0; i < m-n; i++ {
		var k = i
		for j := 0; j < n && haystack[k] == needle[j]; {
			k++
			j++
		}
		if k-i == n {
			return i
		}
	}
	return -1
}
