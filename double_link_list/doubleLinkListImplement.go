package doublelinklist

type MyLinkedList struct {
	Length int
	Head   *ListNode
}

type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if this.Length <= index {
		return -1
	}

	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Length, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length {
		return
	}

	cur := &ListNode{Val: val}

	if index <= 0 {
		cur.Next = this.Head
		if this.Head != nil {
			this.Head.Prev = cur
		}
		this.Head = cur
		this.Length++
		return
	}

	if index == this.Length {
		tmp := this.Head
		for tmp != nil && tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = cur
		cur.Prev = tmp
		this.Length++
		return
	}

	prev := this.Head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	cur.Next = prev.Next
	prev.Next.Prev = cur

	prev.Next = cur
	cur.Prev = prev

	this.Length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Length {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		if this.Head != nil {
			this.Head.Prev = nil
		}

		this.Length--
		return
	}

	prev := this.Head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	delNode := prev.Next
	prev.Next = delNode.Next
	delNode.Next = nil
	delNode.Prev = nil

	if prev.Next != nil {
		prev.Next.Prev = prev
	}

	this.Length--
}
