package go_code

import (
	"math"
)

// https://leetcode.cn/problems/palindrome-number/

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

// 例如，121 是回文，而 123 不是。

// 示例 1：

// 输入：x = 121
// 输出：true
// 示例 2：

// 输入：x = -121
// 输出：false
// 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 示例 3：

// 输入：x = 10
// 输出：false
// 解释：从右向左读, 为 01 。因此它不是一个回文数。

// 提示：

// -231 <= x <= 231 - 1

// 进阶：你能不将整数转为字符串来解决这个问题吗？

func isPalindrome(x int) bool {
	// -1 最小位为 0 必然不是
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	return isPalindromeIntVersionV2(x)
}

func isPalindromeStringVersion(x string) bool {
	i, j := 0, len(x)-1
	for i < j {
		if x[i] != x[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindromeIntVersion(x int) bool {
	var s []int
	var length int
	tmpX := x
	// 取出x的每位的数值 + 位数
	for tmpX > 0 {
		tmp := tmpX % 10
		s = append(s, tmp)
		tmpX /= 10
		length++
	}
	// 反转相乘
	var res int
	for i := 0; i < length; i++ {
		res += s[i] * int(math.Pow10(length-i-1))
	}
	return res == x
}

func isPalindromeIntVersionV2(x int) bool {
	tmpX := x
	res := 0
	for tmpX > 0 {
		// 取x的最后一位 相加 后直接*10，不需要获取位数后再计算
		res = res*10 + tmpX%10 // 123  0*10+3 3*10+2 32*10+1
		tmpX /= 10
	}
	return res == x
}

func isPalindromeIntVersionV3(x int) bool {
	res := 0
	// 只需要反转后半截即可 1234 43 != 12
	// x < res 即代表已经到了到达原始x的一半  1234  123 4；12 43；
	for x > res {
		// 取x的最后一位 相加 后直接*10，不需要获取位数后再计算
		res = res*10 + x%10 // 123  0*10+3 3*10+2 32*10+1
		x /= 10
	}
	// 如果x是奇数位，12345 1234 5；123 54； 12 543；
	return res == x || res/10 == x
}
