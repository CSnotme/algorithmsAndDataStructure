package link_list

// 单链表实现

type MyLinkedList struct {
	size int
	head *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		size: 0,
		head: NewNode(0),
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func GetValues(head *ListNode) []int {
	cur := head

	var values = []int{}
	for cur != nil {
		values = append(values, cur.Val)
		cur = cur.Next
	}

	return values
}

func NewListNodeBySlice(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	var node *ListNode
	var tail *ListNode

	for _, i := range values {
		if tail == nil {
			node = NewNode(i)
			tail = node
		} else {
			tail.Next = NewNode(i)
			tail = tail.Next
		}
	}

	return node
}

// get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1
func (ll *MyLinkedList) Get(index int) int {
	if index < 0 || index >= ll.size {
		return -1
	}

	cur := ll.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

// addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (ll *MyLinkedList) AddAtHead(val int) {
	ll.AddAtIndex(0, val)
}

// addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素
func (ll *MyLinkedList) AddAtTail(val int) {
	ll.AddAtIndex(ll.size, val)
}

// addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，
// 则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。

func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	if index > ll.size {
		return
	}

	if index < 0 {
		index = 0
	}

	newNode := NewNode(val)

	cur := ll.head

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	newNode.Next = cur.Next
	cur.Next = newNode

	ll.size += 1
}

// deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
func (ll *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= ll.size {
		return
	}

	pre := ll.head

	for i := 0; i < index; i++ {
		pre = pre.Next
	}

	pre.Next = pre.Next.Next

	ll.size -= 1

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
