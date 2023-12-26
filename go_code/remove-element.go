package go_code

// https://leetcode.cn/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	head, tail := 0, len(nums)-1
	for head < tail {
		if nums[head] == val {
			nums[head] = nums[tail]
			tail--
		} else {
			head++
		}
	}
	return head
}
