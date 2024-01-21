package go_code

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}
	l, r := 0, 0
	strMap := make(map[byte]int)
	for ; r < len(s); r++ {
		if index, ok := strMap[s[r]]; ok {
			l = max(l, index+1)
		}
		strMap[s[r]] = r
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
