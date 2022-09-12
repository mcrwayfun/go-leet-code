package main

// time complexity: O(m*n)
// space complexity: O(1)
func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return -1
	}

	var m = len(haystack)
	var n = len(needle)
	for i := 0; i <= m-n; i++ {
		var z = i
		var k = 0
		for ; k < n && haystack[z] == needle[k]; {
			z++
			k++
		}
		if k == len(needle) {
			return i
		}
	}
	return -1
}
