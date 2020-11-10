package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var ret []byte

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || c != strs[j][i] {
				return string(ret)
			}
		}
		ret = append(ret, c)
	}

	return string(ret)
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	commonPrefix := longestCommonPrefix(strs)
	fmt.Println(commonPrefix)

	strs = []string{"dog", "racecar", "car"}
	commonPrefix = longestCommonPrefix(strs)
	fmt.Println(commonPrefix)

	strs = []string{"", "racecar"}
	commonPrefix = longestCommonPrefix(strs)
	fmt.Println(commonPrefix)

	strs = []string{"", ""}
	commonPrefix = longestCommonPrefix(strs)
	fmt.Println(commonPrefix)

	strs = []string{"ab", "a"}
	commonPrefix = longestCommonPrefix(strs)
	fmt.Println(commonPrefix)

	strs = []string{"ab", "abb", "a"}
	commonPrefix = longestCommonPrefix(strs)
	fmt.Println(commonPrefix)
}
