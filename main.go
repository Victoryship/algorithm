package main

import (
	"fmt"

	"algorithm/tree"
)

func main() {
	s := []string{"3", "9", "20", "#", "8", "15", "7", "#", "#", "2", "4"}
	root := tree.BuildTreeByArr(s, 0)

	res := tree.PostOrderRecursive(root)
	res1 := tree.PostOrderIteration(root)
	fmt.Println(res)
	fmt.Println(res1)
}
