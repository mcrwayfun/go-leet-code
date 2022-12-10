package main

// time complexity: O(n)
// space complexity: O(n)
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	var stack []byte
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else if len(stack) == 0 {
			return false
		} else {
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if char == '}' && left != '{' {
				return false
			}
			if char == ')' && left != '(' {
				return false
			}
			if char == ']' && left != '[' {
				return false
			}
		}
	}

	return len(stack) == 0
}
