package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Add(val int) *ListNode {
	node := &ListNode{
		Val: val,
	}

	if l == nil {
		return node
	}
	l.Next = node

	return l
}

func (l *ListNode) PrintVal() []int {
	res := []int{}
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func BuildListByArr(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	head := new(ListNode)
	node := head
	for _, v := range data {
		node = node.Add(v)
		node = node.Next
	}

	return head.Next
}
