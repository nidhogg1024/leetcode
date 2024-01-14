package go_code

// https://leetcode.cn/problems/symmetric-tree/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var sameTree func(p *TreeNode, q *TreeNode) bool
	sameTree = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return sameTree(p.Left, q.Right) && sameTree(p.Right, q.Left)
	}
	return sameTree(root.Left, root.Right)
}
