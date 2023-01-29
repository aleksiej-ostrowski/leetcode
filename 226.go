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
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func exchangeLeftAndRight(root *TreeNode) {
	if root == nil {
		return
	}

	temp := root.Right
	root.Right = root.Left
	root.Left = temp

	exchangeLeftAndRight(root.Right)
	exchangeLeftAndRight(root.Left)
}

func invertTree(root *TreeNode) *TreeNode {
	exchangeLeftAndRight(root)
    return root
}

func main() {
	tree := &TreeNode{0, nil, nil}
	res := invertTree(tree)
	fmt.Println(res)
}
