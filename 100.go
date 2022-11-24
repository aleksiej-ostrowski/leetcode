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
	// "sync"
)

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
// from #94
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

/*
func isSameTree2(p *TreeNode, q *TreeNode) bool {

    var pi []int
    var qi []int

    var wg sync.WaitGroup

    wg.Add(2)
    go func(tree *TreeNode, arr *[]int) {
        defer wg.Done()
        *arr = inorderTraversal(tree)
    }(p, &pi)

    go func(tree *TreeNode, arr *[]int) {
        defer wg.Done()
        *arr = inorderTraversal(tree)
    }(q, &qi)

    wg.Wait()

    if len(pi) != len(qi) {
        return false
    }

    for i := 0; i < len(pi); i++ {
        if pi[i] != qi[i] {
            return false
        }
    }

    return true
}
*/

func main() {
	tree1 := &TreeNode{0, nil, nil}
	tree2 := &TreeNode{0, nil, nil}
	res := isSameTree(tree1, tree2)
	fmt.Println(res)
}
