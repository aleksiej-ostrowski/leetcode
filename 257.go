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
	"strings"
)

/**
 * Definition for a binary tree node.
**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePathsHelp(root *TreeNode, path []string) []string {
	if root != nil {
		if (root.Left == nil) && (root.Right == nil) {

			ad := strconv.Itoa(root.Val)
			path = append(path, ad)

			res := strings.Join(path, "->")
			return []string{res}
		}
	} else {
		return []string{}
	}

	ad := strconv.Itoa(root.Val)
	path = append(path, ad)

	c1 := binaryTreePathsHelp(root.Left, path)
	c2 := binaryTreePathsHelp(root.Right, path)

	return append(c1, c2...)
}

func binaryTreePaths(root *TreeNode) []string {
	var path []string
	return binaryTreePathsHelp(root, path)
}

func main() {
	t3 := &TreeNode{8, nil, nil}
	t1 := &TreeNode{5, nil, t3}
	t2 := &TreeNode{6, nil, nil}
	tree := &TreeNode{7, t1, t2}
	r := binaryTreePaths(tree)
	fmt.Println(r)
}
