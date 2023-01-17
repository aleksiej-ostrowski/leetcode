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

func reverseList(head *ListNode) *ListNode {

	if (head == nil) {
		return nil
	}

	x := head
    var new_x *ListNode

	for {
		if x == nil {
			break
		}
        new_x = &ListNode{x.Val, new_x}
        x = x.Next
	}

	return new_x 
}

func main() {
	a := []int{1, 2, 3, 4}
	printList(reverseList(ListCreate(a)))
	fmt.Println("ok")
}
