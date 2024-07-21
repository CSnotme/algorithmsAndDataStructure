package link_list

// 链表找环
// ListNode 在 linkListImplement.go中定义

// 快慢针找环
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

// 快慢针找环的节点
// 找到环并返回环的起点(入环点)
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	// 找相遇点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 如果没找到相遇点
	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

// map找环
func hasCycleByMap(head *ListNode) bool {
	var m = make(map[*ListNode]bool)

	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}

		m[head] = true

		head = head.Next
	}

	return false
}

// 递归删除head找环
// 依次删除head节点，让节点的next指向自己
// 如果有环，那某个节点的next一定指向被删除的某个节点，那这个next在进行递归时，发现自己=自己的next，即有环
func hasCycleByDeleteHeadNode(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	if head.Next == head {
		return true
	}

	nextNode := head.Next

	head.Next = head

	return hasCycleByDeleteHeadNode(nextNode)
}
