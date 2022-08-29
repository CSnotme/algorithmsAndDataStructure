package leetcode

// leetcode 第11题  盛最多水的容器   https://leetcode.cn/problems/container-with-most-water/

// 题目：
/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
*/

// 思路
// 双指针， 左右个一个， 然后计算两节点内的容积，该容积与max容积比较取最大
// 较小的节点移动,左小往右移，右小往左移，直到左右相等停止计算，返回max容积
func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	var tmpArea int
	for left < right {

		if height[right] < height[left] {
			tmpArea = (right - left) * height[right]
			right--
		} else {
			tmpArea = (right - left) * height[left]
			left++
		}

		if tmpArea > maxArea {
			maxArea = tmpArea
		}
	}

	return maxArea

}
