package leetcode

// leetcode 第82题  删除排序链表中的重复元素 II   https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/

// 题目：
/*
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

 示例 1：
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

示例 2：
输入：head = [1,1,1,2,3]
输出：[2,3]

提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列
*/

// 思路
/*
给定的链表是排好序的，因此重复的元素在链表中出现的位置是连续的，因此我们只需要对链表进行一次遍历，就可以删除重复的元素。
由于链表的头节点可能会被删除，因此我们需要额外使用一个哑节点（dummy node）指向链表的头节点。
从指针cur指向链表的哑节点，随后开始对链表进行遍历。
如果当前cur.next 与cur.next.next 对应的元素相同，那么我们就需要将cur.next 以及所有后面拥有相同元素值的链表节点全部删除。
记录一个这个元素的值x，不断将cur.next移除，直到cur.next的值为空或者不等于x。
如果当前cur.next 与cur.next.next对应的元素不相同，说明链表中只有一个元素值为cur.next的节点，那么我们就可以将cur指向cur.next。
*/

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}

	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			tmpVal := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == tmpVal {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next

}
