package link_list

// 链表删除倒数第N个节点

// 双指针
// 定义两个指针从head开始走， 快针先走
// 当快针走到n步时， 慢针开始走
// 当快针走到尾时，慢针就走到倒数第n个的前序位置了
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	fast := head
	slow := head

	// 快针先走n步，
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 如果快针为空，表示链表总长度=n， 要删除的节点是链表头节点
	if fast == nil {
		return head.Next
	}

	// 快慢针开始一起走
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
