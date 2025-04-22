package list

/*
*
 206. 反转链表
    给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
    输入：head = [1,2,3,4,5]
    输出：[5,4,3,2,1]
*/
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p *ListNode
	for head != nil {
		temp := head
		head = head.Next
		temp.Next = p
		p = temp
	}

	return p
}

func DfsReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := DfsReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return p
}

/*
 160. 相交链表
    给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
    图示两个链表在节点 c1 开始相交：
*/
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 != nil && p2 != nil {
			p1 = p1.Next
			p2 = p2.Next
			continue
		}

		if p1 == nil {
			p1 = headB
		}

		if p2 == nil {
			p2 = headA
		}
	}

	return p1
}

/*
 234. 回文链表
    给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

    输入：head = [1,2,2,1]
    输出：true

    输入：head = [1,2]
    输出：false
*/
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 快慢指针找到中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = ReverseList(slow)

	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}

	return true
}

/*
*
 141. 环形链表
    给你一个链表的头节点 head ，判断链表中是否有环。
*/
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

/*
*
*
 142. 环形链表II
    给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
*/
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	// 找到环的入口
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	// 慢指针和头节点移动。相交的第一个几点就是环的入口
	for slow != nil {
		if slow == head {
			return slow
		}
		slow = slow.Next
		head = head.Next
	}

	return nil
}

/*
*

 21. 合并两个有序链表
    将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

    输入：l1 = [1,2,4], l2 = [1,3,4]
    输出：[1,1,2,3,4,4]
*/
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	res := new(ListNode)
	p := res
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	}

	if list2 != nil {
		p.Next = list2
	}

	return res.Next
}

/*
 2. 两数相加
    给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
    请你将两个数相加，并以相同形式返回一个表示和的链表。
    你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

    输入：l1 = [2,4,3], l2 = [5,6,4]
    输出：[7,0,8]
    解释：342 + 465 = 807

    输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
    输出：[8,9,9,9,0,0,0,1]
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		l1 = l2
		l2 = nil
	}

	p := l1
	carry := 0
	for l1 != nil {
		sum := carry + l1.Val
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry, l1.Val = sum/10, sum%10

		if l1.Next == nil && l2 != nil {
			l1.Next = l2
			l2 = nil
		}

		if l1.Next == nil && carry > 0 {
			l1.Next = &ListNode{Val: carry}
			break
		}

		l1 = l1.Next
	}
	return p
}

/*
*

 19. 删除链表的倒数第 N 个结点
    给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

    输入：head = [1,2,3,4,5], n = 2
    输出：[1,2,3,5]

    输入：head = [1], n = 1
    输出：[]
*/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 0 {
		if fast == nil {
			return head
		}

		fast = fast.Next
		n--
	}

	// 链表长度和n相等删除第一个节点
	if fast == nil {
		return head.Next
	}

	for fast != nil {
		if fast.Next == nil {
			slow.Next = slow.Next.Next
			break
		}
		slow = slow.Next
		fast = fast.Next
	}

	return head
}

/*
*
 24. 两两交换链表中的节点
    给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
    输入：head = [1,2,3,4]
    输出：[2,1,4,3]
*/
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := &ListNode{Next: head}
	pre := res
	for pre.Next != nil && pre.Next.Next != nil {
		slow, fast := pre.Next, pre.Next.Next
		slow.Next = fast.Next
		fast.Next = slow
		pre.Next = fast
		pre = slow
	}

	return res.Next
}

func SwapPairsDfs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head.Next
	head.Next = SwapPairsDfs(node.Next)
	node.Next = head
	return node
}

func reverse(head *ListNode, k int) *ListNode {
	var p *ListNode
	start := head
	for head != nil {
		temp := head
		head = head.Next
		temp.Next = p
		p = temp
		k--
		if k == 0 {
			break
		}
	}
	start.Next = head

	return p
}
