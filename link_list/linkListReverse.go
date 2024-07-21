package link_list

// 反转链表， 循环方法
func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 反转链表，递归方法
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
