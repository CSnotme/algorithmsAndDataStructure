package tree

// 树的宽度遍历和深度遍历

// 宽度遍历
// 需要一个栈，先进先出的栈
// 先将头结点压栈
// While栈不为空，
// 		pop一个节点，打印，如果该节点有孩子，先左后右压栈。