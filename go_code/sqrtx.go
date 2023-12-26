package go_code

// https://leetcode.cn/problems/sqrtx/description/

func mySqrt(x int) int {
	res := 1
	for ; res*res <= x; res++ {
	}
	return res - 1
}

func mySqrtV2(x int) int {
	left, right := 0, x
	for left < right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
