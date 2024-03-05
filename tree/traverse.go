package tree

import (
	"strconv"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func BuildTreeByArr(data []string, index int) *TreeNode {
	if len(data) == 0 || index >= len(data) || data[index] == "#" {
		return nil
	}

	n, err := strconv.ParseInt(data[index], 10, 64)
	if err != nil {
		panic(err)
	}

	root := &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   int(n),
	}

	root.Left = BuildTreeByArr(data, 2*index+1)
	root.Right = BuildTreeByArr(data, 2*index+2)

	return root
}

func PreOrderRecursive(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}

	res = append(res, root.Val)
	res = append(res, PreOrderRecursive(root.Left)...)
	res = append(res, PreOrderRecursive(root.Right)...)
	return
}

func PreOrderIteration(root *TreeNode) []int {
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

		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return res
}

func InOrderRecursive(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	res = append(res, InOrderRecursive(root.Left)...)
	res = append(res, root.Val)
	res = append(res, InOrderRecursive(root.Right)...)
	return
}

func InOrderIteration(root *TreeNode) []int {
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

		res = append(res, stack[len(stack)-1].Val)
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return res
}

func PostOrderRecursive(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	res = append(res, PostOrderRecursive(root.Left)...)
	res = append(res, PostOrderRecursive(root.Right)...)
	res = append(res, root.Val)
	return
}

func PostOrderIteration(root *TreeNode) []int {
	var (
		stack = make([]*TreeNode, 0, 16)
		res   = make([]int, 0, 16)
		pre   *TreeNode // 判断右子树是否已入过栈
	)

	for root != nil || len(stack) > 0 {
		// 左子树入栈
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}

		// 右子树入栈
		node := stack[len(stack)-1]
		if node.Right != nil && node.Right != pre {
			root = node.Right
			continue
		}

		// 出栈加入到结果集
		res = append(res, node.Val)
		pre = node
		stack = stack[:len(stack)-1]
	}

	return res
}
