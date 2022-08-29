package leetcode

// leetcode 第3题  无重复字符的最长子串  https://leetcode.cn/problems/longest-substring-without-repeating-characters/

// 题目：
/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/

// 思路： 滑动窗口，用两个指针表示窗口的左右边界， 不停寻找符合要求的情况，并更新结果

// 模板
/*
找最长：
	初始化 left、right、result、bestResult
	while(right没走到结尾){
		while(不满足条件) {
			窗口缩小，移除left对应元素，left右移
		}

		窗口扩大，记录当前right对应元素， 更新最优结果， right右移
	}

找最短：
	初始化 left、right、result、bestResult
	while(right没走到结尾){
		while(满足条件) {
			更新最优结果
			窗口缩小，移除left对应元素，left右移
		}

		窗口扩大，记录当前right对应元素，right右移
	}
*/

func LengthOfLongestSubstring(s string) int {
	maxLen := 0
	curMap := make(map[uint8]int)
	left := 0
	right := 0

	for right < len(s) {
		for curMap[s[right]] != 0 {
			delete(curMap, s[left])
			left++
		}

		curMap[s[right]] = 1
		if maxLen < len(curMap) {
			maxLen = len(curMap)
		}
		right++
	}

	return maxLen
}
