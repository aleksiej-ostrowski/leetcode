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
    Val int
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


func deleteDuplicates(head *ListNode) *ListNode {

    if (head == nil) {
        return nil
    }

    var newL = &ListNode{head.Val, nil}

    var start = newL

    x := head.Next

    for {

        if x == nil {
            break
        }

        if newL.Val != x.Val {
            newL.Next = &ListNode{x.Val, nil}
            newL = newL.Next
        }

        x = x.Next
    }

    return start
}

func main() {

    a := []int{1,1,1,1,1,2,2,3,3,3,3}

    l := ListCreate(a)

    printList(l)

    new_a := deleteDuplicates(l)

    printList(new_a)
    printList(l)

	fmt.Println("ok")
}
