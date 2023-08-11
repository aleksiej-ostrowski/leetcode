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

func averageOfLevels(root *TreeNode) []float64 {

	crn := make(map[int][]float64)

	type Thelper_averageOfLevels func(root *TreeNode, level int)
	var helper_averageOfLevels Thelper_averageOfLevels

	helper_averageOfLevels = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		crn[level] = append(crn[level], float64(root.Val))
		helper_averageOfLevels(root.Left, level+1)
		helper_averageOfLevels(root.Right, level+1)
	}

	helper_averageOfLevels(root, 0)
	otv := make([]float64, len(crn))
	for level, el := range crn {
		len_el := float64(len(el))
		if len_el < .5 {
			continue
		}
		sm := float64(0)
		for _, e := range el {
			sm += e
		}
		otv[level] = sm / len_el
	}
	return otv
}

func main() {
	tree4 := &TreeNode{15, nil, nil}
	tree5 := &TreeNode{7, nil, nil}
	tree3 := &TreeNode{20, tree4, tree5}
	tree2 := &TreeNode{9, nil, nil}
	tree := &TreeNode{3, tree2, tree3}
	fmt.Println(averageOfLevels(tree))
	fmt.Println("ok")
}
