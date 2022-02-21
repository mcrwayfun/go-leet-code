package main

func isValid(s string) bool {
	// 这道题空串也是有效的...
	if len(s) % 2 != 0 {
		return false
	}

	symbol := map[byte]byte {
		')':'(',
		'}':'{',
		']':'[',
	}

	var stack []byte
	for i:=0;i<len(s);i++ {
		if v, ok := symbol[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false // 括号不匹配
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
