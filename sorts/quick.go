package sorts

func QuickSort(data []int, start, end int) {
	if start >= end {
		return
	}

	p := pivot1(data, start, end)
	QuickSort(data, start, p-1)
	QuickSort(data, p+1, end)
}

func pivot(data []int, start, end int) int {
	i, j := start+1, end
	for {
		for i < end && data[i] <= data[start] {
			i++
		}
		for j > start && data[j] >= data[start] {
			j--
		}
		if i >= j {
			break
		}
		data[i], data[j] = data[j], data[i]
	}
	data[start], data[j] = data[j], data[start]

	return j
}

func pivot1(nums []int, start, end int) int {
	i, j := start, end-1
	for {
		for i < end && nums[i] <= nums[end] {
			i++
		}
		for j > start && nums[j] >= nums[end] {
			j--
		}

		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[i], nums[end] = nums[end], nums[i]
	return i
}
