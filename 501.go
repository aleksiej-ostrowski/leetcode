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
    "os"
)

/**
 * Definition for a binary tree node.
**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func moda(data []int) []int {
	// fmt.Println("data = ", data)
	memory := make(map[int]int)
	for _, el := range data {
		memory[el] += 1
	}
	res := []int{0}
	fq := 0
	for key, el := range memory {
		if el > fq {
			fq = el
            res = []int{key}
		} else if el == fq {
			res = append(res, key)
		}
	}
	return res
}

func binaryTreePathsHelp2(root *TreeNode, path []int) []int {
	if root != nil {
		if (root.Left == nil) && (root.Right == nil) {
			path = append(path, root.Val)
			return path
		}
	} else {
		return []int{}
	}

	c1 := binaryTreePathsHelp2(root.Left, path)
	c2 := binaryTreePathsHelp2(root.Right, path)

	res := append(c1, c2...)
	res = append(res, root.Val)

	return res
}

func findMode(root *TreeNode) []int {
	var path []int
	return moda(binaryTreePathsHelp2(root, path))
}

func main() {
    // fmt.Println(moda([]int{0,2,6,4,2,7,9,8,6}))
    // os.Exit(1)
	t3 := &TreeNode{7, nil, nil}
	t1 := &TreeNode{5, nil, t3}
	t2 := &TreeNode{5, nil, nil}
	tree := &TreeNode{7, t1, t2}
	r := findMode(tree)
	fmt.Println(r)
}
