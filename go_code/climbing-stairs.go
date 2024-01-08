package go_code

// https://leetcode.cn/problems/climbing-stairs/description/

func climbStairs(n int) int {
	x, y, z := 0, 0, 1
	for i := 1; i <= n; i++ {
		x, y = y, z
		z = x + y
	}
	return z
}
