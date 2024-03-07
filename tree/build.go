package tree

import "strconv"

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
