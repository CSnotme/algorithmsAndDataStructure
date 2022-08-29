/**
 * @author: caoyongfei(caoyongfei@baidu.com)
 * @file: linkListIntersection
 * @date: 2022-04-25 17:21:36
 * @desc:
**/

package link_list

// 链表相交问题

// 拼接L1+L2  L2+L1的方式 让两个要循环的链表的长度相等
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	n1 := headA
	n2 := headB

	for n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next

		// n1和n2同时为空时， 表示 LA和LB等长， 或者LA+LB 和 LB+LA 等长
		// 如果LA和LB等长时。 走到这里结束， 不会在走下面的交叉遍历
		// 如果LA不等长于LB。 下面就走LA+LB 和 LB+LA
		// 这里满足条件，则没有交点
		if n1 == nil && n2 == nil {
			return nil
		}

		// LA完了，再循环LB
		if n1 == nil {
			n1 = headB
		}

		// LB完了，再循环LA
		if n2 == nil {
			n2 = headA
		}
	}

	return n1
}
