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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode.com/problems/diameter-of-binary-tree/solutions/3362822/java-100-faster-detailed-explaination/

func diameterOfBinaryTreeHelp(root *TreeNode, d_max *int) int {

	if root == nil {
		return 0
	}
    
	max_l := diameterOfBinaryTreeHelp(root.Left, d_max)
	max_r := diameterOfBinaryTreeHelp(root.Right, d_max)

	*d_max = Max(*d_max, max_l + max_r)

	return Max(max_l, max_r) + 1

}

func diameterOfBinaryTree(root *TreeNode) int {
    memory := 0
    _ = diameterOfBinaryTreeHelp(root, &memory)
    return memory
}

func main() {
    /*
		tree3 := &TreeNode{3, nil, nil}
		tree4 := &TreeNode{4, nil, nil}
		tree5 := &TreeNode{5, nil, nil}
		tree2 := &TreeNode{2, tree4, tree5}
		tree1 := &TreeNode{1, tree2, tree3}
    */    
	tree1 := &TreeNode{1, nil, nil}
	fmt.Println(diameterOfBinaryTree(tree1))
	fmt.Println("ok")
}
