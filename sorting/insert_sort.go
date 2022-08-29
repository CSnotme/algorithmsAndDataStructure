package sorting

// 插入排序

/*
1.将数据划分为有序和无序区，开始时0索引有序，1-n无序
2.每次看无序区的一个元素
3.从有序区的最后一个元素开始倒序依次和该元素比大小
4.直到该元素在有序取处于 前面元素 < 该元素 < 后面元素
*/

func InsertSort(nums []int) []int {
	// 0位置已经有序， 从1位置开始，每次看一个元素
	for i := 1; i < len(nums); i++ {
		// i位置往前是已经排序的区域， 当前j对应值依次与前面的值比较， 直到j<=0或者j对应值大于j-1对应值则停止
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	return nums
}
