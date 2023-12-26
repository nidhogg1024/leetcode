package go_code

// https://leetcode.cn/problems/valid-parentheses/description/

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
//
// 示例 1：
//
// 输入：s = "()"
// 输出：true
// 示例 2：
//
// 输入：s = "()[]{}"
// 输出：true
// 示例 3：
//
// 输入：s = "(]"
// 输出：false
//
//
// 提示：
//
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成

func isValid(s string) bool {
	stack := make([]byte, 0, 10)
	for index := range s {
		switch s[index] {
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false

			} else {
				stack = stack[:len(stack)-1]
			}
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s[index])
		}
	}
	return len(stack) == 0
}
