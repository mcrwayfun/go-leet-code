package main

/**
这是一道非常典型的位运算题目，现假设有数字a=9, b=11
两者转换为2进制相加的过程是：
a	1001
b	1011
=  10100

如果只考虑a和b相加，不考虑进位的情况，过程如下：
0 0 = 0
0 1 = 1
1 0 = 1
1 1 = 0
可以得出的结论是，a和b相加，就是a^b的结果

如果只考虑a和b相加后的进位，过程如下：
0 0 = 0
0 1 = 0
1 0 = 0
1 1 = 1
可以得出的结论是，a和b相加的进位，就是a&b的结果

所以a和b相加的过程可以为
1：sum = a^b
2：carry = a&b << 1 (a&b是得到当前位，需要左移得到进位carry）
由于carry不断左移，会在高位补充0，所以最后carry会为0，所以结束的条件便是carry==0
 */
func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := a & b << 1

		a = sum
		b = carry
	}
	return a
}