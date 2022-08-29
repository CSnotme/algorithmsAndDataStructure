/**
 * @author: caoyongfei(caoyongfei@baidu.com)
 * @file: linkListRemoveElement
 * @date: 2022-04-29 17:15:34
 * @desc:
**/

package link_list

// 移除链表元素

// 递归
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = removeElements(head.Next, val)

	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}



// 循环遍历
func removeElementsByFor(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val:  0, Next: nil}
	dummyHead.Next = head
	temp := dummyHead
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return dummyHead.Next
}
