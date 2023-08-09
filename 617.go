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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	v := 0
	if root1 != nil {
		v += root1.Val
	}

	if root2 != nil {
		v += root2.Val
	}

	if root1 != nil || root2 != nil {
		var l1, l2, r1, r2 *TreeNode

		if root1 != nil {
			l1 = root1.Left
			r1 = root1.Right
		}

		if root2 != nil {
			l2 = root2.Left
			r2 = root2.Right
		}

		return &TreeNode{v, mergeTrees(l1, l2), mergeTrees(r1, r2)}
	}

	return nil
}

func main() {
	tree_1_2 := &TreeNode{2, nil, nil}
	tree_1_5 := &TreeNode{5, nil, nil}
	tree_1_3 := &TreeNode{3, tree_1_5, nil}
	tree_1 := &TreeNode{1, tree_1_3, tree_1_2}

	tree_2_4 := &TreeNode{4, nil, nil}
	tree_2_7 := &TreeNode{7, nil, nil}
	tree_2_1 := &TreeNode{1, nil, tree_2_4}
	tree_2_3 := &TreeNode{3, nil, tree_2_7}
	tree_2 := &TreeNode{2, tree_2_1, tree_2_3}

	fmt.Println(mergeTrees(tree_1, tree_2))
	fmt.Println("ok")
}
