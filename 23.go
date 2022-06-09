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
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

const MaxUint = ^uint(0) 
// const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
// const MinInt = -MaxInt - 1

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

    if len(lists) == 0 {
        return nil
    }    

    stop := 0
    for _, e := range lists {
        if e == nil {
            stop += 1
        }
    }

    if stop == len(lists) {
        return nil
    }

    res := &ListNode{0, nil}
    ret := res
    var prv *ListNode = nil

    for {

        r := MaxInt

        for _, e := range lists {
            if e == nil {
                continue
            }
            x := e.Val
            r = min(r, x)
        } 

        // fmt.Println("r = ", r)

        if r == MaxInt {
            if prv != nil {
                prv.Next = nil
            } 
            return ret
        }

        stop := 0

        for i, e := range lists {

            if e == nil {
                stop += 1
                continue
            }

            x := e.Val

            if x == r {
                prv = res
                res.Val = r
                res.Next = &ListNode{0, nil}
                res = res.Next
                lists[i] = e.Next
            }
        }

        if stop == len(lists) {
            if prv != nil {
                prv.Next = nil
            } 
            return ret
        }
    }
}

func ListCreate(a []int) *ListNode {

    if len(a) == 0 {
        return nil
    }

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

func main() {

    a := []int{3, 5, 7}
    b := []int{4, 6, 7, 7, 7, 8, 100}
    c := []int{100, 300, 400}

    l1 := ListCreate(a)
    l2 := ListCreate(b)
    l3 := ListCreate(c)

    printList(l1)
    printList(l2)
    printList(l3)

    r := mergeKLists([]*ListNode{l1, l2, l3})

    fmt.Println("----") 

    printList(r)

}


