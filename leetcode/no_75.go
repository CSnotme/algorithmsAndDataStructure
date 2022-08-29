package leetcode

// leetcode 第75题  颜色分类  https://leetcode.cn/problems/sort-colors/

// 题目：
/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库的sort函数的情况下解决这个问题。



示例 1：
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

示例 2：
输入：nums = [2,0,1]
输出：[0,1,2]

提示：
n == nums.length
1 <= n <= 300
nums[i] 为 0、1 或 2

*/

// 思路
// 荷兰国旗， 快排的partition过程， 将数组分成 小于、等于、大于3个区域
func sortColors(nums []int) {
	red := -1
	blue := len(nums)

	i := 0
	for i < blue {
		if nums[i] < 1 {
			red++
			nums[red], nums[i] = nums[i], nums[red]
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			blue--
			nums[blue], nums[i] = nums[i], nums[blue]
		}
	}

}
