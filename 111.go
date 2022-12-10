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


// https://leetcode.com/problems/minimum-depth-of-binary-tree/solutions/36050/my-simple-recursive-python-solution/

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

    if root.Left == nil {
        return 1 + minDepth(root.Right)
    } 
    
    if root.Right == nil {
        return 1 + minDepth(root.Left)
    }

    return 1 + min(minDepth(root.Left), minDepth(root.Right)) 
}

func main() {
	tree := &TreeNode{0, nil, nil}
	res := minDepth(tree)
	fmt.Println(res)
}
