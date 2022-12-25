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

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if (headA == nil) || (headB == nil) {
		return nil
	}

	xA := headA

	memory := make(map[*ListNode]bool)

	for {

		if xA == nil {
			break
		}

		memory[xA] = true

		xA = xA.Next
	}

	xB := headB

	for {

		if xB == nil {
			break
		}

		if _, ok := memory[xB]; ok {
			return xB
		}

		xB = xB.Next
	}

	return nil
}

func main() {

	a := []int{1, 1, 2}
	b := []int{1, 1, 2}

	printList(getIntersectionNode(ListCreate(a), ListCreate(b)))

	fmt.Println("ok")
}
