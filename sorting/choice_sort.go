package sorting

// 选择排序

/*
每次从未排序数组中找到一个最小值，放到已排序数组中
*/

func ChoiceSort(nums []int) []int {
	// i表示已排序位置， 从0开始， 每次取i开始到最后的最小值， 换到i的位置上， 然后i++
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}


