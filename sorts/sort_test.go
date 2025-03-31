package sorts

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	testCases := [][]int{
		{64, 25, 12, 22, 11},
		{5, 3, 8, 4, 4, 4, 4, 2},
		{1, 2, 3, 4, 5},
		{},
		{1},
		{9, 7, 5, 11, 12, 2, 14, 3, 10, 6},
	}

	for _, testCase := range testCases {
		fmt.Printf("Original: %v\n", testCase)
		sorted := SelectSort(testCase)
		fmt.Printf("Sorted:   %v\n\n", sorted)
	}
}

func TestBubbleSort(t *testing.T) {
	testCases := [][]int{
		{64, 25, 12, 22, 11},
		{5, 3, 8, 4, 2},
		{1, 2, 3, 4, 5},
		{},
		{1},
		{9, 7, 5, 11, 12, 2, 14, 3, 10, 6},
	}

	for _, testCase := range testCases {
		fmt.Printf("Original: %v\n", testCase)
		sorted := BubbleSort(testCase)
		fmt.Printf("Sorted:   %v\n\n", sorted)
	}
}

func TestInsertSort(t *testing.T) {
	testCases := [][]int{
		{64, 25, 12, 22, 11},
		{5, 3, 8, 4, 2},
		{1, 2, 3, 4, 5},
		{},
		{1},
		{9, 7, 5, 11, 12, 2, 14, 3, 10, 6},
	}

	for _, testCase := range testCases {
		fmt.Printf("Original: %v\n", testCase)
		sorted := InsertSort(testCase)
		fmt.Printf("Sorted:   %v\n\n", sorted)
	}
}

func TestQuickSort(t *testing.T) {
	testCases := [][]int{
		{64, 25, 12, 22, 11},
		{5, 3, 8, 4, 2},
		{1, 2, 3, 4, 5},
		{},
		{1},
		{9, 7, 5, 11, 12, 2, 14, 3, 10, 6},
	}

	for _, testCase := range testCases {
		fmt.Printf("Original: %v\n", testCase)
		QuickSort(testCase, 0, len(testCase)-1)
		fmt.Printf("Sorted:   %v\n\n", testCase)
	}
}
