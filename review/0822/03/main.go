package main

// time complexity: O(n)
// space complexity: O(1)
func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p) == 0 {
		return []int{}
	}

	var charFrequencyMap = make(map[byte]int, 256)
	for i := 0; i < len(p); i++ {
		charFrequencyMap[p[i]]++
	}

	var start int
	var matched int
	var res []int
	for end := 0; end < len(s); end++ {
		if _, ok := charFrequencyMap[s[end]]; ok {
			charFrequencyMap[s[end]]--
			if charFrequencyMap[s[end]] == 0 {
				matched++
			}
		}

		if matched == len(charFrequencyMap) {
			res = append(res, start)
		}

		if end - start+1 >= len(p) {
			if _, ok := charFrequencyMap[s[start]]; ok {
				if charFrequencyMap[s[start]] == 0 {
					matched--
				}
				charFrequencyMap[s[start]]++
			}
			start++
		}
	}

	return res
}
