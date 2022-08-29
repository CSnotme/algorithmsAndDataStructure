package leetcode

// leetcode 第147题  对链表进行插入排序  https://leetcode.cn/problems/insertion-sort-list/

// 题目：
/*
给定单个链表的头 head ，使用 插入排序 对链表进行排序，并返回 排序后链表的头 。

插入排序 算法的步骤:

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。
下面是插入排序算法的一个图形示例。部分排序的列表(黑色)最初只包含列表中的第一个元素。每次迭代时，从输入数据中删除一个元素(红色)，并就地插入已排序的列表中。

对链表进行插入排序。

*/

// 思路
// 插入排序的思想， 因为是单链表，因此每次插入我们从链表头开始往后比较
func InsertionSortList(head *ListNode) *ListNode {

	// head节点已经有序，从head的下一个的链表中开始， 每次拿一个
	cur := head.Next
	head.Next = nil
	// 一个假头， 指向head
	newHead := &ListNode{Val: 0, Next: head}

	for cur != nil {
		node := cur
		cur = cur.Next
		node.Next = nil

		prev := newHead
		for {
			if prev != nil && prev.Next != nil {
				if node.Val >= prev.Next.Val {
					prev = prev.Next
				} else {
					node.Next = prev.Next
					prev.Next = node
					break
				}
			} else {
				prev.Next = node
				break
			}
		}
	}

	return newHead.Next
}
