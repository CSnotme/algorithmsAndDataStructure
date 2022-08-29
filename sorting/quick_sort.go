package sorting

import (
	"math/rand"
	"time"
)

// 快排

/*
1.随机选出数组的一个数字(选最右侧)作为划分值，然后根据该数字将数组划分为小于区域、等于区域、大于区域(partition过程)
2.等于区域原地不动，将小于区域和大于区域分别再次进行partition，直到递归结束
3.为保证选出划分值的随机性，最大可能避免快排的最坏情况，每次快排时都随机选择一个元素与队尾交换，使用
*/

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}

	// 随机找数组中一个元素与最后一位交换， 保证划分值的随机性， 最大可能避免快排的最坏情况
	randomIdx := left + RandIntnByTimeSeed(right-left+1)
	nums[randomIdx], nums[right] = nums[right], nums[randomIdx]

	subLeft, subRight := partition(nums, left, right)
	quickSort(nums, left, subLeft-1)
	quickSort(nums, subRight+1, right)
}

// 对 arr[left, rigth]所在区域做划分
// 取arr[right]为划分值， 将区域划分成，左边小于划分值，中间等于划分值，右边大于划分值
// 返回划分后等于区域的左右边界
func partition(nums []int, left int, right int) (int, int) {
	less := left - 1 // 小于区的右边界
	more := right    // 大于区的左边界

	// 从区域最左侧下标依次往后看， 直到左侧下标等于大于区域的左边界停止
	for left < more {
		if nums[left] < nums[right] {
			// 如果小于划分值
			// 小于区的右边界++
			less++
			// 交换小于区的右边界元素和left所对应元素， (第一次交换是相同的值)
			nums[left], nums[less] = nums[less], nums[left]
			// 左侧下标++(看下个元素)
			left++
		} else if nums[left] > nums[right] {
			// 如果大于划分值
			// 大于区左边界--
			more--
			// 交换大于区左边界元素和left所对应元素
			nums[left], nums[more] = nums[more], nums[left]
			// 不需要left++,因为交换后的元素还没看呢
		} else {
			// 这是等于划分值，原地不动, left++
			left++
		}
	}

	// 最后交换大于区左边界和划分值
	nums[more], nums[right] = nums[right], nums[more]

	return less + 1, more
}

// 使用时间戳作为种子随机生成0-N之间的数字
func RandIntnByTimeSeed(n int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(n)
}
