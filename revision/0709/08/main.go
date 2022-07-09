package main

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	var stack []byte
	for i:=0; i<len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 {
			return false
		} else {
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if s[i] == ')' && right != '(' {
				return false
			}
			if s[i] == '}' && right != '{' {
				return false
			}
			if s[i] == ']' && right != '[' {
				return false
			}
		}
	}

	return len(stack) == 0
}
