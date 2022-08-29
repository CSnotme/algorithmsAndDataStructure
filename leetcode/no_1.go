package leetcode

// leetcode 第1题  两数之和   https://leetcode.cn/problems/two-sum/

// 题目：
/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。.

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]


提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
*/

// 思路
// 搞个map把所有val和index的映射记录一下
// 我们拿到一个数之后，另一个数等于target-这个数， 我们通过map查找另一个数是否存在即可
func TwoSum(nums []int, target int) []int {

	valIdxMap := make(map[int]int)

	// 把val和index映射到map中
	for idx, i := range nums {
		valIdxMap[i] = idx
	}

	// 遍历
	for idx, i := range nums {
		// 要找的数
		ret := target - i
		// 判断是否在map中，并且对应的index不是原数
		if retIdx, ok := valIdxMap[ret]; ok && idx != retIdx {
			return []int{idx, retIdx}
		}
	}

	return []int{}

}
