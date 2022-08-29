package tree

// 是否完全二叉树

// 使用宽度遍历，每个节点不满足下面任意一个条件的都不是完全二叉树
// 1.节点有右孩子没左孩子
// 2.首次出现左右孩子不满时，后面的节点都应该没孩子
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type message struct {
	Node  *TreeNode
	Level int
}

type Stack struct {
	Data []message
}

func NewStack() Stack {
	return Stack{
		Data: []message{},
	}
}
func (s *Stack) Push(msg message) {
	s.Data = append(s.Data, msg)
}

func (s *Stack) Pop() message {
	msg := s.Data[0]
	s.Data = s.Data[1:]
	return msg
}

func (s *Stack) Empty() bool {
	return len(s.Data) == 0
}
