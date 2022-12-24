// from 144.go

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

func postorderTraversal_recc(root *TreeNode) []int {    
    if root == nil {    
        return []int{}    
    }    
    var res []int    
    if root.Left != nil {    
        c1 := postorderTraversal_recc(root.Left)    
        res = append(res, c1...)    
    }    
    if root.Right != nil {    
        c2 := postorderTraversal_recc(root.Right)    
        res = append(res, c2...)    
    }    
    res = append(res, root.Val)    
    return res   
}

func postorderTraversal(root *TreeNode) []int {
    return postorderTraversal_recc(root)
}

func main() {
	tree1 := &TreeNode{1, nil, nil}
	tree2 := &TreeNode{2, nil, nil}
	tree := &TreeNode{0, tree1, tree2}
	res := postorderTraversal(tree)
	fmt.Println(res)
}
