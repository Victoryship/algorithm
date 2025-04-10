package tree

/*
104. 二叉树的最大深度
给定一个二叉树 root ，返回其最大深度。
二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
*/
func MaxDepth(root *TreeNode) int {
	var res, depth int
	var traverse func(r *TreeNode)
	traverse = func(r *TreeNode) {
		if r == nil {
			return
		}

		depth++
		if res < depth {
			res = depth
		}
		traverse(r.Left)
		traverse(r.Right)
		depth--
	}

	traverse(root)
	return res
}

// 通过递归实现
func MaxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := MaxDepth1(root.Left)
	r := MaxDepth1(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}

/*
226. 翻转二叉树
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
输入：root = [2,1,3]
输出：[2,3,1]
*/
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}

/*
*

 101. 对称二叉树
    给你一个二叉树的根节点 root ， 检查它是否轴对称。

    输入：root = [1,2,2,3,4,4,3]
    输出：true

    输入：root = [1,2,2,null,3,null,3]
    输出：false

    核心思想
    通过递归比较左子树和右子树的节点是否满足镜像对称条件：

    终止条件：
    如果两个节点都为空 → 对称，返回 true。
    如果其中一个节点为空 → 不对称，返回 false。
    如果两个节点的值不相等 → 不对称，返回 false。
    递归条件：
    比较 左子树的左节点 与 右子树的右节点（外侧）。
    比较 左子树的右节点 与 右子树的左节点（内侧）。
    两组比较结果的 逻辑与 为最终结果。
*/
func IsSymmetricRecursion(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// 迭代法实现
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		l, r := queue[0], queue[1]
		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil || l.Val != r.Val {
			return false
		}

		queue = queue[2:]
		queue = append(queue, l.Left, r.Right)
		queue = append(queue, l.Right, r.Left)
	}

	return true
}

/*
*
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
*/
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}

	Flatten(root.Left)
	Flatten(root.Right)

	temp := root.Right
	root.Left, root.Right = nil, root.Left

	for root.Right != nil {
		root = root.Right
	}
	root.Right = temp
}

// 迭代实现
func Flatten2(root *TreeNode) {
	if root == nil {
		return
	}

	pre := new(TreeNode)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil {
			pre.Left, pre.Right = nil, node
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		pre = node
	}
}

/*
*
543. 二叉树的直径
给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。

输入：root = [1,2,3,4,5]
输出：3
解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
*/
func DiameterOfBinaryTree(root *TreeNode) int {
	var res int
	var depth func(tree *TreeNode) int
	depth = func(tree *TreeNode) int {
		if tree == nil {
			return 0
		}
		l := depth(tree.Left)
		r := depth(tree.Right)
		if res < l+r {
			res = l + r
		}
		if l > r {
			return l + 1
		}

		return r + 1
	}
	depth(root)
	return res
}
