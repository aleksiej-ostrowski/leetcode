package main

import (
    "fmt"
)

/**
 * Definition for a binary tree node.
**/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
// https://leetcode.com/problems/binary-tree-inorder-traversal/solutions/127563/binary-tree-inorder-traversal/

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

func main() {
    tree := &TreeNode{0, nil, nil}
    res := inorderTraversal(tree)
    fmt.Println(res)
}

