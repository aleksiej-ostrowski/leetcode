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

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getMinimumDifferenceHelp(root *TreeNode, pred *[]int, dd int) int {
	if root != nil {
		dd_pred := 10_000_000
		for _, el := range *pred {
			dd_pred = Min(dd_pred, Abs(el-root.Val))
		}
		dd = Min(dd, dd_pred)
		*pred = append(*pred, root.Val)
		if (root.Left == nil) && (root.Right == nil) {
			return dd
		}
	} else {
		return dd
	}

	min_el_l := getMinimumDifferenceHelp(root.Left, pred, dd)
	min_el_r := getMinimumDifferenceHelp(root.Right, pred, dd)

	return Min(min_el_l, min_el_r)

}

func getMinimumDifference(root *TreeNode) int {
	return getMinimumDifferenceHelp(root, &[]int{1_000_000}, 100_000_000)
}

func main() {
	tree4 := &TreeNode{499, nil, nil}
	tree2 := &TreeNode{424, nil, tree4}
	tree5 := &TreeNode{689, nil, nil}
	tree3 := &TreeNode{612, nil, tree5}
	tree1 := &TreeNode{600, tree2, tree3}
	fmt.Println(getMinimumDifference(tree1))
	fmt.Println("ok")
}
