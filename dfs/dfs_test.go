package dfs

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{[]int{2}, [][]int{{2}}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v,%d", tc.nums, tc.want), func(t *testing.T) {
			originalNums := make([]int, len(tc.nums))
			copy(originalNums, tc.nums) // 保存原始数组的副本

			got := Permute(tc.nums)
			fmt.Println("res", got)
		})
	}
}

func TestPermuteUnique(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{[]int{2}, [][]int{{2}}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v,%d", tc.nums, tc.want), func(t *testing.T) {
			originalNums := make([]int, len(tc.nums))
			copy(originalNums, tc.nums) // 保存原始数组的副本

			got := PermuteUnique(tc.nums)
			fmt.Println("res", got)
		})
	}
}
