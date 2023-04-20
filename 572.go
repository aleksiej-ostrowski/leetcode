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

//https://leetcode.com/problems/subtree-of-another-tree/solutions/3349136/java-easy-solution-0ms-beats100/

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val == subRoot.Val {
		if isIdentical(root, subRoot) {
			return true
		}
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isIdentical(node *TreeNode, subRoot *TreeNode) bool {

	if node == nil && subRoot == nil {
		return true
	} else if node == nil || subRoot == nil || node.Val != subRoot.Val {
		return false
	}

	if !isIdentical(node.Left, subRoot.Left) {
		return false
	}

	if !isIdentical(node.Right, subRoot.Right) {
		return false
	}

	return true
}

/*

// It is not working.

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	root_v := inorderTraversal(root)
	sub_v := inorderTraversal(subRoot)
    fmt.Println(root_v)
    fmt.Println(sub_v)
	for idx_m, el_m := range root_v {
		if el_m == sub_v[0] {
			sum := 0
			idx := idx_m
			for _, el_s := range sub_v {
				if idx >= len(root_v) {
					break
				}
				if el_s == root_v[idx] {
					sum += 1
					idx += 1
				}
			}
			if sum == len(sub_v) {
				return true
			}
		}
	}
	return false
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	curr := root
	for {
		if curr == nil {
			break
		}
		if curr.Left == nil {
			res = append(res, curr.Val)
			curr = curr.Right
		} else {
			pre := curr.Left
			for {
				if pre.Right == nil {
					break
				}
				pre = pre.Right
			}
			pre.Right = curr
			// temp := &TreeNode{curr.Val, curr.Left, curr.Right}
			temp := curr
			curr = curr.Left
			temp.Left = nil
		}
	}

	return res
}


*/

func main() {
	/*
		tree3 := &TreeNode{3, nil, nil}
		tree4 := &TreeNode{4, nil, nil}
		tree5 := &TreeNode{5, nil, nil}
		tree2 := &TreeNode{2, tree4, tree5}
		tree1 := &TreeNode{1, tree2, tree3}
	*/
	tree1 := &TreeNode{1, nil, nil}
	tree2 := &TreeNode{1, nil, nil}
	fmt.Println(isSubtree(tree1, tree2))
	fmt.Println("ok")
}
