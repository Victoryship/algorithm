package tree

import (
	"fmt"
	"testing"
)

func TestBuildTreeByArr(t *testing.T) {
	a := BuildTreeByArr([]string{"1", "2", "3", "4", "5", "6", "7"}, 0)
	fmt.Println(PreOrderRecursive(a))
	fmt.Println(InOrderRecursive(a))
	fmt.Println(PostOrderRecursive(a))
}

func TestPreOrderHelper(t *testing.T) {
	root := BuildTreeByArr([]string{"1", "2", "3", "4", "5", "6", "7"}, 0)
	if root == nil {
		return
	}

	var (
		stack = make([]*TreeNode, 0, 16)
		res   = make([]int, 0, 16)
	)

	for root != nil || len(stack) > 0 {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
			continue
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
}

func TestInOrderHelper(t *testing.T) {
	root := BuildTreeByArr([]string{"1", "2", "3", "4", "5", "6", "7"}, 0)
	if root == nil {
		return
	}

	var (
		stack = make([]*TreeNode, 0, 16)
		res   = make([]int, 0, 16)
	)

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}

		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
}

func TestPostOrderHelper(t *testing.T) {
	root := BuildTreeByArr([]string{"1", "2", "3", "4", "5", "6", "7"}, 0)
	if root == nil {
		return
	}

	var (
		stack = make([]*TreeNode, 0, 16)
		res   = make([]int, 0, 16)
		pre   = new(TreeNode) // pre用来判断是否左右节点已经加入过栈中的节点
	)

	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		// 栈中节点的左节点和右节点同时为空，表示遍历到最底层。此时可以把结果加入到返回数组中
		// 或者栈弹出的节点的左节点或右节点就是上一次栈弹出的节点。那么此时弹出的节点就是上次的根节点
		if (node.Left == nil && node.Right == nil) || (pre != nil && pre == node.Left || pre == node.Right) {
			res = append(res, node.Val)
			pre = node
			stack = stack[:len(stack)-1]
			continue
		}

		// 这里要注意先把右节点入栈，这样左节点的输出顺序才会在右节点之前
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}
