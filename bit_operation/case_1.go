package bitoperation

// 位运算相关内容

// 题目1：在一个数组中，1个数出现了奇数次，其他数都出现了偶数次，求出现奇数次的数

// 题目2：在一个数组中，2个数出现了奇数次，其他数都出现了偶数次，求这两个出现奇数次的数

// 解题目1
// 依次做异或运算，因为偶数次的元素异或之后为0， 再异或奇数次元素就等于该元素
func answer_1(nums []int) int {
	ero := 0
	for _, v := range nums {
		ero = ero ^ v
	}

	return ero
}

// 解题目2
// 1.假设这两个数a和b都是出现奇数次，首先所有元素异或之后会得到 eor = a^b
// 2.因为a和b不相等，所以a^b这个值在某一bit上等于1的时候，说明a和b在该bit上的值不相等
// 3.我们可以根据元素该bit位置上否为1， 将整个数组一分为2，那a和b自然分别落在不同的区间
// 4.任意求一个区间的异或得到的结果 eor' 要么是a，要么是b
// 5. eor ^ eor' 就能得到另一个
func Answer_2(nums []int) (int, int) {

	// 所有元素的异或
	ero := 0
	for _, v := range nums {
		ero = ero ^ v
	}

	// ero的bit上取左右侧=1
	rightBitOne := ero & (^ero + 1)

	// 计算a或者b的值
	onlyOne := 0
	for _, v := range nums {
		// 取该bit上为0的元素
		if v&rightBitOne == 0 {
			onlyOne ^= v
		}
	}

	return onlyOne, ero ^ onlyOne
}
