/*

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
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

func sumOfLeftLeaves_recc(root *TreeNode, what int) int {
	if root == nil {
		return 0
	}
	var res int = 0

    if (root.Left == nil) && (root.Right == nil) && (what == 1) {
	    res += root.Val
    }

	if root.Left != nil {
		res += sumOfLeftLeaves_recc(root.Left, 1)
	}
	if root.Right != nil {
		res += sumOfLeftLeaves_recc(root.Right, 2)
	}
	return res
}


func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeaves_recc(root, 0)
}


func main() {
	tree1 := &TreeNode{1, nil, nil}
	tree2 := &TreeNode{2, nil, nil}
	tree := &TreeNode{0, tree1, tree2}
	res := sumOfLeftLeaves(tree)
	fmt.Println(res)
}
