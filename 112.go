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

func search_recc(root *TreeNode, sum int, targetSum int) bool {

    if (root != nil) {
	    if (root.Left == nil) && (root.Right == nil) {
		    return sum + root.Val == targetSum
	    }
    } else {
        return false
    }
   
    val := sum + root.Val

	c1 := search_recc(root.Left, val, targetSum)
	c2 := search_recc(root.Right, val, targetSum)

	return c1 || c2
}


func hasPathSum(root *TreeNode, targetSum int) bool {

	return search_recc(root, 0, targetSum)
}


func main() {
	tree := &TreeNode{0, nil, nil}
	res := hasPathSum(tree, 0)
	fmt.Println(res)
}
