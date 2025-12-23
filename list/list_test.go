package list

import (
	"fmt"
	"testing"
)

func TestBuildListByArr(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	head := BuildListByArr(data)

	if head == nil {
		t.Error("Expected head to be non-nil")
	}

	current := head
	for _, v := range data {
		if current.Val != v {
			t.Errorf("Expected %d, got %d", v, current.Val)
		}
		current = current.Next
	}

	if current != nil {
		t.Error("Expected current to be nil at the end of the list")
	}
}

func TestListNode_PrintVal(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	head := BuildListByArr(data)

	result := head.PrintVal()
	if len(result) != len(data) {
		t.Errorf("Expected length %d, got %d", len(data), len(result))
	}

	for i, v := range data {
		if result[i] != v {
			t.Errorf("Expected %d, got %d", v, result[i])
		}
	}
	fmt.Println(result)
}

func TestReverseList(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	head := BuildListByArr(data)
	reversedHead := ReverseList2(head)

	result := reversedHead.PrintVal()
	expected := []int{5, 4, 3, 2, 1}

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d, got %d", v, result[i])
		}
	}
	fmt.Println(result)
}

func TestMergeTwoLists(t *testing.T) {
	data1 := []int{1, 2, 4}
	data2 := []int{1, 3, 4}
	list1 := BuildListByArr(data1)
	list2 := BuildListByArr(data2)

	mergedHead := MergeTwoLists(list1, list2)

	result := mergedHead.PrintVal()
	expected := []int{1, 1, 2, 3, 4, 4}

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d, got %d", v, result[i])
		}
	}
	fmt.Println(result)
}

func TestReverseKGroup(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	head := BuildListByArr(data)
	prev := new(ListNode)
	prev.Next = head
	temp := prev

	res := reverse(head, 3)
	temp.Next = res
	result := res.PrintVal()
	temp = head
	expected := []int{5, 4, 3, 2, 1}

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	fmt.Println(result)
}
