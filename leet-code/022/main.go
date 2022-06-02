package main

import "fmt"

/**
数字 n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：
func generateParenthesis(n int) []string {}
 */

/**
形成有效的括号组合，什么是有效的？前提是小于n
对于(, 任何情况下都是有效的
对于)，只要有一个(, )就是有效的

所以有规则，假设left和right分别表示可以生成的左右括号的次数
只要有 left>0, 就可以添加左括号
只要有 right>left, 就可以添加右括号

time complexity: O(n!)
space complexity: O(1)
 */
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}

	var res []string
	generate(&res, "", n, n)
	return res
}

func generate(res *[]string, str string, left int, right int) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
	} else {
		if left > 0 {
			generate(res, str+"(", left-1, right)
		}
		if right > left {
			generate(res, str+")", left, right-1)
		}
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}


