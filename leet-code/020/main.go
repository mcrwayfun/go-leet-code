package main

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 1：
输入：s = "()"
输出：true

示例2：
输入：s = "()[]{}"
输出：true

示例3：
输入：s = "(]"
输出：false

示例4：
输入：s = "([)]"
输出：false

示例5：
输入：s = "{[]}"
输出：true

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isValid(s string) bool {}
*/

// time complexity: O(n)
// space complexity: O(n)
func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 {
			return false
		} else {
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i] == ')' && right != '(' {
				return false
			}
			if s[i] == ']' && right != '[' {
				return false
			}
			if s[i] == '}' && right != '{' {
				return false
			}
		}
	}
	return len(stack) == 0
}
