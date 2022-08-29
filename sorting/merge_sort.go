package sorting

// 归并排序

/*
1.将数组一分为2，每个小数组在进行归并排序，然后将两个排好序的小数组合并成有序(merge过程)
*/

func MergeSort(nums []int) []int {
	return process(nums)
}

// 归并排序中的合并两数组
func merge(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}

	if len(arr2) == 0 {
		return arr1
	}

	// 开辟新的空间存放合并后的结果
	newArr := []int{}
	i := 0
	j := 0

	// 从每个数组的第一个元素开始对比，直到某个数据的元素全部完成，跳出循环
	for i < len(arr1) && j < len(arr2) {
		// 比较数组1的i对应元素和数组2的j对应元素
		// i对应的小，将i对应的值加入新数组， i++
		// 否则将j对应的值加入新数组， j++
		if arr1[i] <= arr2[j] {
			newArr = append(newArr, arr1[i])
			i++
		} else {
			newArr = append(newArr, arr2[j])
			j++
		}
	}

	// 跳出循环之后，可能有一个数组还没读完

	// 数组1是否读完
	for i < len(arr1) {
		newArr = append(newArr, arr1[i])
		i++
	}

	// 数组2是否读完
	for j < len(arr2) {
		newArr = append(newArr, arr2[j])
		j++
	}

	return newArr
}

// 递归二分数组并排序
func process(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	leftArr := process(nums[:mid])
	rightArr := process(nums[mid:])

	return merge(leftArr, rightArr)
}
