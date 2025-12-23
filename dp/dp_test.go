package dp

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{5, 4, 1, 2}, 1},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{1, 1, 1, 1}, 0},
		{[]int{2}, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v,%d", tc.nums, tc.want), func(t *testing.T) {
			originalNums := make([]int, len(tc.nums))
			copy(originalNums, tc.nums) // 保存原始数组的副本

			got := Trap(tc.nums)
			if got != tc.want {
				t.Errorf("Trap(%v) = %d, want %d", originalNums, got, tc.want)
			}
		})
	}
}
