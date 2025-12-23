package array

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		str  string
		want string
	}{
		{"babad", "bab"},
		{"", ""},
		{"s", "s"},
		{"sssssssss", "sssssssss"},
		{"cbbd", "bb"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v,%s", tc.str, tc.want), func(t *testing.T) {
			got := LongestPalindrome(tc.str)
			if got != tc.want {
				t.Errorf("LongestPalindrome(%v) = %s, want %s", tc.str, got, tc.want)
			}
		})
	}
}

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{2}, 2},
		{[]int{}, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v,%d", tc.nums, tc.want), func(t *testing.T) {
			originalNums := make([]int, len(tc.nums))
			copy(originalNums, tc.nums) // 保存原始数组的副本

			got := MaxArea(tc.nums)
			if got != tc.want {
				t.Errorf("MaxArea(%v) = %d, want %d", originalNums, got, tc.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums   []int
		want   int
		wantNs []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 1, 2, 2, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1}, 1, []int{1}},
		{[]int{}, 0, []int{}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			originalNums := make([]int, len(tc.nums))
			copy(originalNums, tc.nums) // 保存原始数组的副本

			got := RemoveDuplicates(tc.nums)
			if got != tc.want {
				t.Errorf("RemoveDuplicates(%v) = %d, want %d", originalNums, got, tc.want)
			}

			// 检查新数组的内容
			expected := tc.wantNs
			actual := tc.nums[:got]
			if len(actual) != len(expected) {
				t.Errorf("RemoveDuplicates(%v) resulted in slice %v, expected %v", originalNums, actual, expected)
			} else {
				for i := range actual {
					if actual[i] != expected[i] {
						t.Errorf("RemoveDuplicates(%v) resulted in slice %v, expected %v", originalNums, actual, expected)
						break
					}
				}
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		nums   []int
		val    int
		want   int
		wantNs []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
		{[]int{1, 2, 3, 4, 5}, 6, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1}, 1, 0, []int{}},
		{[]int{}, 1, 0, []int{}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, %d", tc.nums, tc.val), func(t *testing.T) {
			originalNums := make([]int, len(tc.nums))
			copy(originalNums, tc.nums) // 保存原始数组的副本

			got := RemoveElement(tc.nums, tc.val)
			if got != tc.want {
				t.Errorf("RemoveElement(%v, %d) = %d, want %d", originalNums, tc.val, got, tc.want)
			}

			// 检查新数组的内容
			expected := tc.wantNs
			actual := tc.nums[:got]
			if len(actual) != len(expected) {
				t.Errorf("RemoveElement(%v, %d) resulted in slice %v, expected %v", originalNums, tc.val, actual, expected)
			} else {
				for i := range actual {
					if actual[i] != expected[i] {
						t.Errorf("RemoveElement(%v, %d) resulted in slice %v, expected %v", originalNums, tc.val, actual, expected)
						break
					}
				}
			}
		})
	}
}

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
