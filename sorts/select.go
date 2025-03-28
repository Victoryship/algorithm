package sorts

func SelectSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if minIndex == i {
			continue
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	return nums
}
