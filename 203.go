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

/*
func removeElements2(head *ListNode, val int) *ListNode {

	if (head == nil) {
		return nil
	}

	x := head

    var new_x *ListNode
    start := new_x

    nodes := 0

	for {

		if x == nil {
			break
		}

        if x.Val != val {

            if nodes == 0 {
                new_x = &ListNode{0, nil}
                start = new_x
            } else {
                new_x.Next = &ListNode{0, nil}
                new_x = new_x.Next
            }

            new_x.Val = x.Val
            nodes++
        }

		x = x.Next
	}

	return start 
}
*/

func removeElements(head *ListNode, val int) *ListNode {

	if (head == nil) {
		return nil
	}

	x := head
    start := head
    x_old := head

	for {
		if x == nil {
			break
		}
        if x.Val != val {
            x_old = x
        } else {
            x_old.Next = x.Next
        }
        x = x.Next
	}

    if start.Val == val {
        return start.Next
    }

	return start 
}

func main() {
	a := []int{7, 7, 7, 7}
	printList(removeElements(ListCreate(a), 7))
	fmt.Println("ok")
}
