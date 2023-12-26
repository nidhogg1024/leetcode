package go_code

// https://leetcode.cn/problems/plus-one/

func plusOne(digits []int) []int {
	var result []int
	add := 0
	for i := len(digits) - 1; i >= 0; i-- {
		one := 0
		if i == len(digits)-1 {
			one = 1
		}
		value := (digits[i] + one + add) % 10
		add = (digits[i] + one + add) / 10
		result = append(result, value)
	}
	if add != 0 {
		result = append(result, add)
	}
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}

func plusOneV2(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			// 如果是 9 赋值为 0，下一次前一位直接加 1 即可
			digits[i] = 0
		}
	}
	// digits 中所有的元素均为 9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
