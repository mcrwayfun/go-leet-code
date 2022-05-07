package main

/**
实现pow(x, n)，即计算 x 的 n 次幂函数（即，xn ）。

示例 1：
输入：x = 2.00000, n = 10
输出：1024.00000

示例 2：
输入：x = 2.10000, n = 3
输出：9.26100

示例 3：
输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/powx-n
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func myPow(x float64, n int) float64
*/

/**
解题思路：这道题是求x的n次方，可以通过循环n次求得结果.
当n非常大的时候，会超出了时间的限制

time complexity: O(n)
space complexity: O(1)
*/
func myPow(x float64, n int) float64 {
	var result float64 = 1
	var N = abs(n)

	for i := 0; i < N; i++ {
		result *= x
	}

	if n < 0 {
		return 1 / result
	}
	return result
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

/**
解题思路2：假设现在求 2^11次方，则这里x=2，n=11
2^11=2^8 * 2^2 * 2^1, 将11转换成二进制的：1011。
先假设N=11, result=1, x=2
1：N&1==1=true, result*=x=1*2=2, x*=x=2*2=4, N>>1=0101=5
2：N&1==1=true, result*=x=2*4=8, x*=x=4*4=16, N>>1=0010=4
3：N&1==0==false, x*=x=16*16=256, N>>1=0001=2
4：N&1==1==true, result*=x=512, x*=x=256*256=65536, N>>1=0000

这个解题思路简单来说：
1：每次循环会判断N的二进制最后一位是否为1，如果为1就将result*x
2：每次x都会乘以x自己，N都会将N右移一位
3：循环直到 N == 0 结束

time complexity: O(lgn) // 循环变量N每次会变为一半
space complexity: O(1)
*/
func myFastPow(x float64, n int) float64 {
	var result float64 = 1
	var N = abs(n)

	for N != 0 {
		if N&1 == 1 {
			result *= x
		}
		x *= x
		N >>= 1
	}

	if n < 0 {
		return 1 / result
	}
	return result
}
