/**
 * @author: caoyongfei(caoyongfei@baidu.com)
 * @file: linkListOddEven
 * @date: 2022-05-06 15:56:00
 * @desc:
**/

package link_list

// 奇偶链表


// 给定单链表的头节点 head，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
// 第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
// 请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var evenHead = head.Next
	var odd = head
	var even = evenHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}
