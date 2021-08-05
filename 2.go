/*

#---------------------------------------------------#
#                                                   #
#  version 0.0.1                                    #
#  https://leetcode.com/problems/add-two-numbers/   #
#                                                   #
#  Aleksiej Ostrowski, 2020                         #
#                                                   #
#  https://aleksiej.com                             #
#                                                   #
#---------------------------------------------------#

*/

package main

import (
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

/*

func reverseList(l *ListNode) *ListNode {

    temp := []int{}

    x := l

    for {

        if x == nil {
            break
        }

        temp = append(temp, x.Val)

        x = x.Next
    }

    // https://stackoverflow.com/questions/19239449/how-do-i-reverse-an-array-in-go
    for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
        temp[i], temp[j] = temp[j], temp[i]
    }

    var newl = &ListNode{0, nil}

    var start = newl

    for i, e := range temp {
        newl.Val = e
        if i != len(temp)-1 {
            newl.Next = &ListNode{0, nil}
            newl = newl.Next
        }
    }

    return start
}
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    if (l1 == nil) || (l2 == nil) {
        return nil
    }

    x1 := l1
    x2 := l2

    var newL = &ListNode{0, nil}

    var start = newL

    flag1 := true
    flag2 := true

    extra := 0

    for {

        newVal := extra

        if flag1 {
            newVal += x1.Val
        }

        if flag2 {
            newVal += x2.Val
        }

        extra = 0

        if newVal > 9 {
            extra = 1
            newVal -= 10
        }

        newL.Val = newVal

        if flag1 {
            x1 = x1.Next
            flag1 = x1 != nil
        }

        if flag2 {
            x2 = x2.Next
            flag2 = x2 != nil
        }

        if flag1 || flag2 || (extra > 0) {
            newL.Next = &ListNode{0, nil}
            newL = newL.Next
        } else {
            break
        }
    }

    return start

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

func main() {

    a := []int{2, 4, 3}
    b := []int{5, 6, 4}

    l1 := ListCreate(a)
    l2 := ListCreate(b)

    printList(l1)
    printList(l2)

    printList(addTwoNumbers(l1, l2))

    // 7 0 8
}
