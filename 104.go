/*

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #

*/

package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth_recc(root *TreeNode, cnt int) int {

	if root == nil {
		return cnt
	}

	cnt += 1

	c1 := maxDepth_recc(root.Left, cnt)
	c2 := maxDepth_recc(root.Right, cnt)

	if c1 > c2 {
		return c1
	} else {
		return c2
	}

}

func maxDepth(root *TreeNode) int {

	return maxDepth_recc(root, 0)

}

func main() {
	tree := &TreeNode{0, nil, nil}
	res := maxDepth(tree)
	fmt.Println(res)
}
