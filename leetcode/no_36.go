package leetcode

// leetcode 第36题  有效的数独   https://leetcode.cn/problems/valid-sudoku/

// 题目：
/*
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）


注意：
一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。

提示：
board.length == 9
board[i].length == 9
board[i][j] 是一位数字（1-9）或者 '.'

*/

// 思路
// 计数算法， 记录每一行的数字的出现次数， 记录每一列的数字的出现次数， 记录每一个九宫格的数字的出现次数
// 将所有元素遍历一遍， 然后计数， 然后判断是否在行\列\方块中出现超过一次

func isValidSudoku(board [][]byte) bool {
	var rowsList [9][9]int  // 记录每一行的数字的出现次数
	var colsList [9][9]int  // 记录每一列的数字的出现次数
	var sqList [3][3][9]int // 记录每一个九宫格的数字的出现次数

	for i, row := range board {
		for j, c := range row {
			// 跳过空的格子
			if c == '.' {
				continue
			}

			// 计数通过下标0-8记录1-9的出现次数，如rowsList[0][0] = 1 表示第1行数字1出现了1次
			// 因此这里用数字-1计算的是下标
			idx := c - '1'

			rowsList[i][idx]++      // 行计数+1
			colsList[j][idx]++      // 列计数+1
			sqList[i/3][j/3][idx]++ // 九宫格计数+1

			// 判断是否某个计数>1
			if rowsList[i][idx] > 1 || colsList[j][idx] > 1 || sqList[i/3][j/3][idx] > 1 {
				return false
			}

		}
	}

	return true

}
