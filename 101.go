/*

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
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
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil) && (q == nil) {
		return true
	}
	curr1 := p
	curr2 := q
	for {
		if (curr1 == nil) || (curr2 == nil) {
			if curr1 != curr2 {
				return false
			}
			break
		}
		if (curr1.Left == nil) || (curr2.Left == nil) {
			if curr1.Left != curr2.Left {
				return false
			}
			if curr1.Val != curr2.Val {
				return false
			}
			curr1 = curr1.Right
			curr2 = curr2.Right
		} else {
			pre1 := curr1.Left
			pre2 := curr2.Left
			for {
				if (pre1.Right == nil) || (pre2.Right == nil) {
					if pre1.Right != pre2.Right {
						return false
					}
					break
				}
				pre1 = pre1.Right
				pre2 = pre2.Right
			}

			pre1.Right = curr1
			temp1 := curr1
			curr1 = curr1.Left
			temp1.Left = nil

			pre2.Right = curr2
			temp2 := curr2
			curr2 = curr2.Left
			temp2.Left = nil
		}
	}

	return true
}

func exchangeLeftAndRight(root *TreeNode) {
	if root == nil {
		return
	}

	temp := root.Right
	root.Right = root.Left
	root.Left = temp

	exchangeLeftAndRight(root.Right)
	exchangeLeftAndRight(root.Left)
}

func isSymmetric(root *TreeNode) bool {
	a := root.Right
	b := root.Left

	exchangeLeftAndRight(a)

	return isSameTree(a, b)
}

func main() {
	tree := &TreeNode{0, nil, nil}
	res := isSymmetric(tree)
	fmt.Println(res)
}
