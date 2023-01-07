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
 * Definition for singly-linked list.
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListCreate(a []int) *ListNode {

	var l = &ListNode{0, nil}

	var start = l

	for i, s := range a {

		l.Val = s
		if i != len(a)-1 {
			l.Next = &ListNode{0, nil}
			l = l.Next
		}
	}

	return start
}

func printList(l *ListNode) {

	x := l
	for {
		if x == nil {
			break
		}
		fmt.Print(x.Val, " ")
		x = x.Next
	}

	fmt.Println()
}


func removeElements(head *ListNode, val int) *ListNode {

	if (head == nil) {
		return nil
	}

	x := head

    new_x := &ListNode{0, nil}
    start := new_x

	for {

		if x == nil {
			break
		}

        if x.Val != val {
            if new_x == nil {
                new_x.Next = &ListNode{0, nil}
            }
            new_x.Val = x.Val
            new_x = new_x.Next
        }

		x = x.Next
	}

	return start 
}

func main() {
	a := []int{1, 2, 6, 3, 4, 5, 6}
	printList(removeElements(ListCreate(a), 6))
	fmt.Println("ok")
}
