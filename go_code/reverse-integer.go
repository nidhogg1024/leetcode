package go_code

import "math"

func reverse(x int) int {
	var res int
	for x != 0 {
		num := x % 10
		// max int32 是 2147483647
		if res > 214748364 || (res == 214748364 && num > 7) {
			return 0
		}
		// min int32 是 -2147483648
		if res < -214748364 || (res == -214748364 && num < -8) {
			return 0
		}
		res = res*10 + num
		x /= 10
	}
	return res
}

func reverseV2(x int) int {
	var res int
	for x != 0 {
		num := x % 10
		// 不需要判断最后一位，输入的int不会移除，所以最后一位只会是1或2
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		res = res*10 + num
		x /= 10
	}
	return res
}
