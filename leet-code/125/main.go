package main

import "fmt"

/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
示例 2:

输入: "race a car"
输出: false
解释："raceacar" 不是回文串

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

题目模板：
func isPalindrome(s string) bool
*/

/**
题目解析：正着和反转过来都是相同的字符串可以认为是回文串，题目中
1：只考虑字母和数字字符，可以忽略字符的大小写
2：空格也定义为有效的回文串

综上条件可以得出的结论是：
1：无论大小写，可以将字符统一转换为大写或者小写来比较
2：空格是有效回文串，即可以忽略空格

解法：
1：编写辅助函数isAlphanumeric(b byte)，判断是不是字母或者数字
2：编写辅助函数isEqualIgnoreCase(a, b byte)，
	2-1：可以通过 +32 将字符a和b转换为小写（ASCII码中 小写字母=大写字母+32）
	2-2：判断a和b是否相等
3：维护两个指针i和j，i从下标0开始，j从数组末尾开始，开始循环
4：如果i指向的元素为非字母或者数字，则i++，直到指定的元素为字母或者数字
5：如果j指向的元素为非字母或者数字，则j--，直到指定的元素为字母或者数字
6：判断i和j指定的元素是否相等，如果不相等则结束循环并返回false
7：循环结束后都没有找到不相等的情况，则认为是回文串，返回true

Time Complexity: O(n)
Space Complexity: O(1)
*/
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}
		for i < j && !isAlphanumeric(s[j]) {
			j--
		}
		if i < j && !isEqualIgnoreCase(s[i], s[j]) {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}

func isEqualIgnoreCase(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 32
	}
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	return a == b
}

func main() {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	fmt.Println(isPalindrome(s1))
	fmt.Println(isPalindrome(s2))
}
