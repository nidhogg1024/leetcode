package go_code

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		s := haystack[i : i+len(needle)]
		if s == needle {
			return i
		}
	}
	return -1
}
