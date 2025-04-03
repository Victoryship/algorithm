package sorts

func ShellSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}

	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i += gap {
			data := nums[i]
			j := i - gap
			for j >= 0 && nums[j] > data {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = data
		}
	}

	return nums
}
