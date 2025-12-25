package dfs

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		nums string
		want []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"2", []string{"a", "b", "c"}},
		{"9", []string{"w", "x", "y", "z"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v,%v", tc.nums, tc.want), func(t *testing.T) {
			got := LetterCombinations(tc.nums)
			for i := 0; i < len(got); i++ {
				if got[i] != tc.want[i] {
					t.Errorf("LetterCombinations(%s) = %s, want %s", tc.nums, got[i], tc.want[i])
				}
			}
		})
	}
}

func TestLetterCombinations1(t *testing.T) {
	testCases := []struct {
		num int
		want []string
	}{
		{3, []string{"000", "001", "010", "011", "100", "101", "110", "111"}},
		{2, []string{"00", "01", "10", "11"}},
		{1, []string{"0", "1"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d,%v", tc.num, tc.want), func(t *testing.T) {
			got := LetterCombinations1(tc.num)
			for i := 0; i < len(got); i++ {
				if got[i] != tc.want[i] {
					t.Errorf("LetterCombinations(%d) = %s, want %s", tc.num, got[i], tc.want[i])
				}
			}
		})
	}
}

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
