package go_code

// https://leetcode.cn/problems/length-of-last-word/description/

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && res > 0 {
			return res
		} else {
			res++
		}
	}
	return res
}
