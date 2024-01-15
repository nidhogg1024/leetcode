package go_code

// https://leetcode.cn/problems/add-two-numbers/description/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	res := &ListNode{}
	cur := res
	for l1 != nil || l2 != nil {
		var a, b int
		if l1 == nil {
			a = 0
		} else {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			b = 0
		} else {
			b = l2.Val
			l2 = l2.Next
		}
		val := (a + b + carry) % 10
		carry = (a + b + carry) / 10
		cur.Next = &ListNode{
			Val: val,
		}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return res.Next
}
