package main

import (
	"fmt"
)

// Definition for a binary tree node.

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// https://leetcode.com/problems/balanced-binary-tree/solutions/35686/java-solution-based-on-height-check-left-and-right-node-in-every-recursion-to-avoid-further-useless-search/

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func height(node *TreeNode) int {

    if node == nil {
        return 0
    }

    lH := height(node.Left)

    if lH == -1 {
        return -1
    }
    
    rH := height(node.Right)

    if rH == -1 {
        return -1
    }

    if lH - rH < -1 || lH - rH > 1 {
        return -1
    }

    return max(lH, rH) + 1
}

func isBalanced(root *TreeNode) bool {
    if (root == nil) {
        return true
    }
    return height(root) != -1
}

func main() {
	tree := &TreeNode{0, nil, nil}
	res := isBalanced(tree)
	fmt.Println(res)
}
