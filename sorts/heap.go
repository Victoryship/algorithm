package sorts

/*
1、堆是一种完全二叉树
2、对于最大堆：所有的父节点都大于等于它的子节点
3、对于最小堆：所有的父节点都小于等于它的子节点
4、完全二叉树左节点：2*i+1
5、完全二叉树右节点：2*i+2
6、完全二叉树父节点：i/2
7、完全二叉树最后一个非叶子节点：len(nums)/2-1（数组长度除一半减1）
*/
func HeapSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}

	// 构建最大堆
	for i := length/2 - 1; i >= 0; i-- {
		heapify(nums, i, length)
	}

	// 交换堆顶元素和最后一个元素，并调整堆
	for i := length - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i)
	}

	return nums
}

// 交换堆顶元素和最后一个元素，并调整堆
func heapify(nums []int, i, length int) {
	maxIndex := i    // 初始化最大值索引为父节点
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点

	// 比较左子节点是否更大
	if left < length && nums[left] > nums[maxIndex] {
		maxIndex = left
	}

	// 比较右子节点是否更大
	if right < length && nums[right] > nums[maxIndex] {
		maxIndex = right
	}

	// 如果最大值不是父节点，交换并递归调整子树
	if maxIndex != i {
		nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
		heapify(nums, maxIndex, length)
	}
}
