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

https://leetcode.com/problems/binary-tree-tilt/solutions/3219060/best-java-solution-recursive-dfs-java-accepted/

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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findTiltHelp(root *TreeNode, res *int) int {

	if root == nil {
		return 0
	}

	left := findTiltHelp(root.Left, res)
	right := findTiltHelp(root.Right, res)

	*res += Abs(left - right)

	return left + right + root.Val
}

func findTilt(root *TreeNode) int {
	res := 0
	_ = findTiltHelp(root, &res)
	return res
}

func main() {
	tree6 := &TreeNode{7, nil, nil}
	tree5 := &TreeNode{9, nil, tree6}

	tree3 := &TreeNode{3, nil, nil}
	tree4 := &TreeNode{5, nil, nil}

	tree2 := &TreeNode{2, tree3, tree4}
	tree1 := &TreeNode{4, tree2, tree5}
	fmt.Println(findTilt(tree1))
	fmt.Println("ok")
}
