package main

// time complexity: O(n^2)
// space complexity: O(1)
func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return -1
	}

	var m = len(haystack)
	var n = len(needle)
	for i := 0; i <= m-n; i++ {
		var j = i
		var z = 0
		for ; z < n && haystack[j] == needle[z] ; {
			j++
			z++
		}
		if z == len(needle) {
			return i
		}
 	}
 	return -1
}
