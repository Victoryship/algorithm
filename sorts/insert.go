package sorts

func InsertSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}

	for i := 1; i < length; i++ {
		data := nums[i]
		j := i-1
		for j >= 0 && nums[j] > data {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1]  = data // 插入位置为j+1，因为多减去了1
	}

	return nums
}