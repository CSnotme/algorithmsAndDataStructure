package tree

import "math"

// 判断是否是搜索二叉树

// 搜索二叉树， 即每个子树满足 左孩子<根<右孩子
// 中序遍历后，应该得到一个升序列表

var prevVal = math.MinInt

func IsBst(head *Node) bool {
	if head == nil {
		return true
	}

	// 判断左子树是否bst
	if !IsBst(head.Left) {
		return false
	}

	// 判断当前节点值是否小于上一个值
	if head.Val < prevVal {
		return false
	} else {
		prevVal = head.Val
	}

	// 右子树是否bst
	return IsBst(head.Right)
}
