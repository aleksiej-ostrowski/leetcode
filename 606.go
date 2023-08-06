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
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := strconv.Itoa(root.Val)
	left := root.Left
	right := root.Right
	if left != nil {
		res += "(" + tree2str(left) + ")"
	}
	if right != nil {
		if left == nil {
			res += "()"
		}
		res += "(" + tree2str(right) + ")"
	}
	return res
}

func main() {
	tree4 := &TreeNode{4, nil, nil}
	tree3 := &TreeNode{3, nil, nil}
	tree2 := &TreeNode{2, tree4, nil}
	tree := &TreeNode{1, tree2, tree3}
	fmt.Println(tree2str(tree))
	fmt.Println("ok")
}
