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

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	x := head.Next

	memory := make(map[*ListNode]bool)

	for {

		if x == nil {
			break
		}

		if _, ok := memory[x]; ok {
			return true
		}

		memory[x] = true

		x = x.Next
	}

	return false
}

func main() {

	a := []int{1, 1, 1, 1, 1, 2, 2, 3, 3, 3, 3}

	l := ListCreate(a)

	printList(l)

	fmt.Println(hasCycle(l))

	fmt.Println("ok")
}
