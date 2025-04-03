package tree

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
			// 仅当节点存在子节点时才入栈
			if root.Left != nil || root.Right != nil {
				stack = append(stack, root)
			}
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

		// 右子树先入栈
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
