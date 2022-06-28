package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}

	var n = len(haystack)
	var m = len(needle)
	for i:=0; i<=n-m; i++ {
		j := 0
		z := i
		for ;j<m && haystack[z] == needle[j]; {
			j++
			z++
		}

		if j == len(needle) {
			return i
		}
	}
	return -1
}
