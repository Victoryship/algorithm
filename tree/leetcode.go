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

/*
*
105. 从前序与中序遍历序列构造二叉树
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
*/
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	index := 0
	for k, v := range inorder {
		if v == preorder[0] {
			index = k
			break
		}
	}

	root.Left = BuildTree(preorder[1:index+1], inorder[:index])
	root.Right = BuildTree(preorder[index+1:], inorder[index+1:])
	return root
}

/*
*
98. 验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。

输入：root = [3,1,5,0,2,4,6]
输出：false
*/
func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(*TreeNode, int, int) bool
	dfs = func(tree *TreeNode, minVal, maxVal int) bool {
		if tree == nil {
			return true
		}

		if tree.Val <= minVal || tree.Val >= maxVal {
			return false
		}
		return dfs(tree.Left, minVal, tree.Val) && dfs(tree.Right, tree.Val, maxVal)
	}
	return dfs(root, -1<<32, 1<<32)
}

/*
*
230. 二叉搜索树中第 K 小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。
输入：root = [3,1,4,null,2], k = 1
输出：1
*/
func KthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var (
		stack = []*TreeNode{root}
	)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}

		node := stack[len(stack)-1]
		k--
		if k == 0 {
			return node.Val
		}
		root = node.Right
		stack = stack[:len(stack)-1]
	}

	return 0
}

/*
*
236. 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
*/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	l := LowestCommonAncestor(root.Left, p, q)
	r := LowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}

	return r

}

/*
*
199. 二叉树的右视图
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

输入：root = [1,2,3,null,5,null,4]
输出：[1,3,4]
*/
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var (
		res = []int{}
		dfs func(r *TreeNode, depth int)
	)

	dfs = func(r *TreeNode, depth int) {
		if r == nil {
			return
		}

		if depth > len(res) {
			res = append(res, r.Val)
		}

		dfs(r.Right, depth+1)
		dfs(r.Left, depth+1)
	}

	dfs(root, 1)

	return res
}

/**
124. 二叉树中的最大路径和
	二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
	路径和 是路径中各节点值的总和。
	给你一个二叉树的根节点 root ，返回其 最大路径和 。

	输入：root = [-10,9,20,null,null,15,7]
	输出：42
	解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
*/

func MaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		maxValue = -10000
		dfs      func(r *TreeNode) int
	)
	dfs = func(tree *TreeNode) int {
		if tree == nil {
			return 0
		}
		l := dfs(tree.Left)
		r := dfs(tree.Right)
		if l < 0 {
			l = 0
		}
		if r < 0 {
			r = 0
		}
		if tree.Val+l+r > maxValue {
			maxValue = tree.Val + l + r
		}

		if l > r {
			return tree.Val + l
		}

		return tree.Val + r
	}

	dfs(root)
	return maxValue
}

/*
*

 437. 路径总和 III
    给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
    路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

    输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
    输出：3
    解释：和等于 8 的路径有 3 条

    前缀和记录路径条数
*/
func PathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	prefix := make(map[int]int, 32)
	prefix[0] = 1
	var dfs func(tree *TreeNode, sum int) int
	dfs = func(tree *TreeNode, sum int) int {
		if tree == nil {
			return 0
		}
		sum += tree.Val
		count := prefix[sum-targetSum]
		prefix[sum]++
		count += dfs(tree.Left, sum)
		count += dfs(tree.Right, sum)
		prefix[sum]--
		return count
	}

	return dfs(root, 0)
}
