// from 94.go

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

// https://leetcode.com/problems/binary-tree-inorder-traversal/solutions/127563/binary-tree-inorder-traversal/

/*

func preorderTraversal(root *TreeNode) []int {
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

func preorderTraversal_recc(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
    var res []int
    res = append(res, root.Val)
	if root.Left != nil {
        c1 := preorderTraversal_recc(root.Left)
		res = append(res, c1...)
	}
	if root.Right != nil {
        c2 := preorderTraversal_recc(root.Right)
		res = append(res, c2...)
	}
	return res
}

func preorderTraversal(root *TreeNode) []int {
	return preorderTraversal_recc(root)
}

func main() {
	tree1 := &TreeNode{1, nil, nil}
	tree2 := &TreeNode{2, nil, nil}
	tree := &TreeNode{0, tree1, tree2}
	res := preorderTraversal(tree)
	fmt.Println(res)
}
