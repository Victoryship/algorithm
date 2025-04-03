package sorts

func MergeSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}

	mid := length / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	lengthL, lengthR := len(left), len(right)
	res := make([]int, 0, lengthL+lengthR)
	i, j := 0, 0

	for i < lengthL && j < lengthR {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
			continue
		}
		res = append(res, right[j])
		j++
	}
	// 可能还有剩余元素未加入到res中
	if i < lengthL {
		res = append(res, left[i:]...)
	}

	if j < lengthR {
		res = append(res, right[j:]...)
	}

	return res
}
