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
	"sort"
)

/**
 * Definition for a binary tree node.
**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	keys := make(map[int]int)
	findSecondMinimumValue_help(root, keys)
	res := make([]int, 0, len(keys))
	for key := range keys {
		res = append(res, key)
	}
	keys = nil
	sort.Ints(res)
	x0 := res[0]
	for _, el := range res[1:] {
		if x0 != el {
			return el
		}
	}
	return -1
}

func findSecondMinimumValue_help(root *TreeNode, keys map[int]int) {

	if root == nil {
		return
	}

	keys[root.Val] = 1

	findSecondMinimumValue_help(root.Left, keys)
	findSecondMinimumValue_help(root.Right, keys)
}

func main() {
	tree3 := TreeNode{5, nil, nil}
	tree4 := TreeNode{7, nil, nil}
	tree2 := TreeNode{5, &tree3, &tree4}
	tree1 := TreeNode{2, nil, nil}
	tree0 := TreeNode{2, &tree1, &tree2}

	fmt.Println(findSecondMinimumValue(&tree0))

	fmt.Println("ok")
}
