package go_code

// https://leetcode.cn/problems/check-if-a-string-is-an-acronym-of-words/description/?envType=daily-question&envId=2023-12-20

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i := 0; i < len(words); i++ {
		if words[i][0] != s[i] {
			return false
		}
	}
	return true
}
