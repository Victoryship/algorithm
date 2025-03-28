package sorts

func BubbleSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := 0; i < n; i++ {
		isSort := false
		for j := n - 1; j > i; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				isSort = true
			}
		}
		if !isSort {
			break
		}
	}

	return nums
}
