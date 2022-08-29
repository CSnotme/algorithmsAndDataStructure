/**
 * @author: caoyongfei(caoyongfei@baidu.com)
 * @file: linkListPalindrome
 * @date: 2022-05-06 16:36:29
 * @desc:
**/

package link_list

import "fmt"

// 回文链表

func isPalindrome(head *ListNode) (bool, *ListNode) {
	isP := true

	slow := head
	fast := head.Next


	// 找链表中点
	// 慢走1 ，快走2， 快走完时， 慢走到中间(奇数个元素链表), 慢走到中间2元素的左边元素(偶数个元素链表)
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 后半段的链表反转
	var prev *ListNode
	cur := slow.Next
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	// 遍历两个链表比较
	left := head
	right := prev
	for left != nil && right != nil {
		if left.Val != right.Val {
			isP = false
			break
		}
		left = left.Next
		right = right.Next
	}

	// 恢复链表结构
	cur = prev
	prev = nil
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	slow.Next = prev

	return isP, head
}

func Test1(head *ListNode) {
	fmt.Println(GetValues(head))
	b, newHead := isPalindrome(head)
	fmt.Println(b)
	fmt.Println(GetValues(newHead))
}

