package main

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return -1
	}

	var n = len(haystack)
	var m = len(needle)
	for i := 0; i <= n-m; i++ {
		var k = i
		var j = 0
		for ; haystack[k] == needle[j]; {
			k++
			j++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
