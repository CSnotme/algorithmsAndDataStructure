package leetcode

// leetcode 第19题  删除链表的倒数第 N 个结点  https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
// 题目：
/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

提示：
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
*/

// 双指针
// 定义两个指针从head开始走， 快针先走
// 当快针走到n步时， 慢针开始走
// 当快针走到尾时，慢针就走到倒数第n个的前序位置了
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	delPre := head

	// 快针先走n步，
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 快针走到最后了， 说明要删除的是头结点
	if fast == nil {
		return head.Next
	}
	// 快慢针开始一起走
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		delPre = delPre.Next
	}

	delPre.Next = delPre.Next.Next
	return head
}
