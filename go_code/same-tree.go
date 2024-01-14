package go_code

// https://leetcode.cn/problems/same-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTreeBfs(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	queueP := []*TreeNode{p}
	queueQ := []*TreeNode{q}

	for len(queueP) > 0 && len(queueQ) > 0 {
		curP, curQ := queueP[0], queueQ[0]
		queueP, queueQ = queueP[1:], queueQ[1:] // 出队列
		if curP.Val != curQ.Val {
			return false
		}
		if (curP.Left == nil && curQ.Left != nil) || (curP.Left != nil && curQ.Left == nil) {
			return false
		}
		if (curP.Right == nil && curQ.Right != nil) || (curP.Right != nil && curQ.Right == nil) {
			return false
		}
		if curP.Left != nil {
			queueP = append(queueP, curP.Left)
		}
		if curP.Right != nil {
			queueP = append(queueP, curP.Right)
		}
		if curQ.Left != nil {
			queueQ = append(queueQ, curQ.Left)
		}
		if curQ.Right != nil {
			queueQ = append(queueQ, curQ.Right)
		}
	}
	return len(queueP) == len(queueQ)
}
