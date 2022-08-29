package sorting

// 堆排序

// i的左右孩子节点  2*i+1 2*i+2
// i的父节点   (i - 1) / 2

func HeapSort(nums []int) []int {
	if nums == nil || len(nums) < 2 {
		return nums
	}

	// 先把初始的数组变成大根堆
	// 从下标0开始， 依次对数组做heapInsert， 因为heapInset只会对idx前的值生效，就相当于每次增加一个新元素到堆底，做一次heapInsert
	for i := 0; i < len(nums); i++ {
		heapInsert(nums, i)
	}

	heapSize := len(nums)

	// 弹出堆顶最大元素， 交换最后元素
	nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
	// 然后heapSize--缩小堆的范围，即交换到尾部的弹出元素不再属于堆，而是属于排序数组
	heapSize--

	// 当堆不为空时
	for heapSize > 0 {
		// 每次交换的元素都会放在堆顶，做heapify，变成大根堆
		heapify(nums, 0, heapSize)
		// 弹出堆顶最大元素， 交换最后元素
		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
		// 然后heapSize--缩小堆的范围，即交换到尾部的弹出元素不再属于堆，而是属于排序数组
		heapSize--
	}

	return nums
}

// 大根堆中，由堆中某索引位置开始， 在堆中做向上运动， 将堆调整为大根堆
func heapInsert(nums []int, idx int) {
	// 如果该索引位置的元素比他的父节点大， 就交换
	for nums[idx] > nums[(idx-1)/2] {
		nums[idx], nums[(idx-1)/2] = nums[(idx-1)/2], nums[idx]
		idx = (idx - 1) / 2
	}
}

// 由堆某索引位置开始，在堆中向下做运动, 将堆调整为大根堆
func heapify(nums []int, idx int, heapSize int) {
	// 左孩子的下标
	leftIdx := idx*2 + 1

	// 只要还有孩子就循环
	for leftIdx < heapSize {
		// 获取两个孩子中值最大的孩子的下标
		// 最大下标先用左孩子
		maxIdx := leftIdx
		// 如果有右孩子， 并且右孩子的值大于左孩子的值， 用右孩子的下标
		if leftIdx+1 < heapSize && nums[leftIdx+1] > nums[leftIdx] {
			maxIdx = leftIdx + 1
		}

		// 如果孩子的最大值大于当前节点， 就交换
		if nums[maxIdx] > nums[idx] {
			nums[maxIdx], nums[idx] = nums[idx], nums[maxIdx]
			// 更新idx和leftIdx
			idx = maxIdx
			leftIdx = idx*2 + 1
		} else {
			// 否则退出循环
			break
		}
	}
}
