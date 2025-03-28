package sorts

func QuickSort(data []int, start, end int) {
	if start >= end {
		return
	}

	p := pivot(data, start, end)
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
