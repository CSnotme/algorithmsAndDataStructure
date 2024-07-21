package leetcode

// leetcode 第946题  验证栈序列  https://leetcode.cn/problems/validate-stack-sequences/

// 题目：
/*
给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，
只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。

示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

*/

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	poppedIdx := 0

	for _, i := range pushed {
		// push元素入栈
		stack = append(stack, i)

		// 比较栈顶元素 和 poppedInx所指向的元素
		// 如果相等，栈顶出栈，poppedIdx+1
		for len(stack) > 0 && stack[len(stack)-1] == popped[poppedIdx] {
			stack = stack[:len(stack)-1]
			poppedIdx++
		}
	}
	return len(stack) == 0
}
