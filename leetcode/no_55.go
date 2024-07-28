package leetcode

// leetcode 第55题  跳跃游戏  https://leetcode.cn/problems/jump-game/description/

// 题目：
/*
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。


示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


提示：

1 <= nums.length <= 104
0 <= nums[i] <= 105

*/

/*
思路：如下，
1.计算每个索引最大可跳的位置，maxJump = nums[i] + i
2.遍历索引<=maxJump的位置，不停地更新最大可跳到的索引位置 maxJump
3.判断maxJump是否可跳到最后索引
索引：0,1,2,3,4
步长：2,3,4,5,1
最远：2,4,6,8,5

*/

func canJump(nums []int) bool {

	maxJump := nums[0] + 0
	// 遍历每个元素，
	for i := 0; i < len(nums); i++ {
		if maxJump >= i {
			if nums[i]+i > maxJump {
				maxJump = nums[i] + i
			}

			if maxJump >= len(nums)-1 {
				return true
			}
		} else {
			return false
		}
	}

	return false
}
