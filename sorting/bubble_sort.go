package sorting

// 冒泡排序

/*
每次都从未排序区数组0索引元素开始，依次比较相邻两数的大小，将较大的的数交换到最后，然后该数划分到已排序区
*/

func BubbleSort(nums []int) []int {
	// 遍历数组每一个元素，数组长度-i-1表示已经排序的元素
	for i := 0; i < len(nums)-1-i; i++ {
		// 每次从0索引开始到未排序区域的结尾，进行冒泡比较
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}

		}
	}

	return nums
}

