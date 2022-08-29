package tree

import "fmt"

// 二叉树遍历

// 递归法的基础结构， 递归序， 在递归序不同的位置进行打印，即可以实现先序、中序、后序的遍历
func TraverseTree(head *Node) {
	if head == nil {
		return
	}
	TraverseTree(head.Left)
	TraverseTree(head.Right)
}

// 递归先序遍历。 头、左、右
// 先序非递归的遍历，需要一个栈，
// 先把头节点压栈，
// while栈不为空:
// 		就pop一个节点，打印该节点，
// 		压入右节点, 压入左节点
func TraverseTreeFirst(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	TraverseTreeFirst(head.Left)
	TraverseTreeFirst(head.Right)
}

// 递归中序遍历。 左、头、右
// 中序非递归遍历， 需要一个栈。
// 1.从该结点开始左边界依次压栈，直到左边界为Null
// 2.pop一个节点，打印，如果有右节点，重复1步
func TraverseTreeMiddle(head *Node) {
	if head == nil {
		return
	}
	TraverseTreeMiddle(head.Left)
	fmt.Println(head.Val)
	TraverseTreeMiddle(head.Right)
}

// 递归后序遍历。 左、右、头
// 中序非递归遍历， 需要两个栈。 栈1，栈2，
// 栈1弹出顺序为头右左，弹出一个就压入栈2，栈2弹出顺序即头右左的逆序，左右头
// 头节点压入栈1
// 1.每次pop一个节点， 压入栈2
// 2.如果该节点有孩子节点， 先压左，再压右 到栈1
// 重复1,2步， 直到栈1为空
// 依次pop栈2并打印
func TraverseTreeLast(head *Node) {
	if head == nil {
		return
	}
	TraverseTreeLast(head.Left)
	TraverseTreeLast(head.Right)
	fmt.Println(head.Val)
}
