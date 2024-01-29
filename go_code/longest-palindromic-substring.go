package go_code

// https://leetcode.cn/problems/longest-palindromic-substring/description/

// 状态转移方程 dp[i][j]  = dp[i+1][j-1] || s[i] == s[j]
// 边界条件 (j-1) - (i + 1) + 1 < 2 简化 j - i < 3
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	maxLen, begin := 1, 0
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 { // 表示去除首尾两个字符，长度为0或1
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}
