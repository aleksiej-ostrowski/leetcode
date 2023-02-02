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
	"os"
)

type MyStack struct {
	stack []int
}

func ConstructorS() MyStack {
	stack := make([]int, 0)
	return MyStack{stack}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() (bool, int) {
	if this.Empty() {
		return false, 0
	}
	len_stack := len(this.stack)
	el := this.stack[len_stack-1]
	this.stack = this.stack[:len_stack-1]
	return true, el
}

func (this *MyStack) Top() (bool, int) {
	if this.Empty() {
		return false, 0
	}
	len_stack := len(this.stack)
	el := this.stack[len_stack-1]
	return true, el
}

func (this *MyStack) Empty() bool {
	return this.Size() == 0
}

func (this *MyStack) Size() int {
	return len(this.stack)
}

type MyQueue struct {
	s1    MyStack
	s2    MyStack
	current int
}

func Constructor() MyQueue {
	var queue MyQueue
	queue.s1 = ConstructorS()
	queue.s2 = ConstructorS()
	return queue
}

func (this *MyQueue) Push(x int) {
    this.s1.Push(x)
}

func (this *MyQueue) Peek() int {

    for {
        ok, el := this.s1.Pop()
        if ok {
            this.s2.Push(el)
        } else {
            break
        }
    }

    _, res := this.s2.Top()

    for {
        ok, el := this.s2.Pop()
        if ok {
            this.s1.Push(el)
        } else {
            break
        }
    }

    return res
}


func (this *MyQueue) Pop() int {

    for {
        ok, el := this.s1.Pop()
        if ok {
            this.s2.Push(el)
        } else {
            break
        }
    }

    _, res := this.s2.Pop()

    for {
        ok, el := this.s2.Pop()
        if ok {
            this.s1.Push(el)
        } else {
            break
        }
    }

    return res
}

func (this *MyQueue) Size() int {
	return len(this.s1.stack)
}

func (this *MyQueue) Empty() bool {
	return this.Size() == 0
}

func (this *MyQueue) Print() {

    for {
        ok, el := this.s1.Pop()
        if ok {
            this.s2.Push(el)
        } else {
            break
        }
    }

    for {
        ok, el := this.s2.Pop()
        if ok {
            this.s1.Push(el)
            fmt.Println("el = ", el)
        } else {
            break
        }
    }
}

func main() {
	/*
	   obj := ConstructorS()
	   obj.Push(10)
	   obj.Push(20)
	   obj.Push(30)
	   fmt.Println(obj.stack)

	   ok, el1 := obj.Pop()
	   if ok {
	       fmt.Println("pop = ", el1)
	   }

	   fmt.Println(obj.stack)

	   obj.Push(40)

	   ok, el1 = obj.Pop()
	   if ok {
	       fmt.Println("pop = ", el1)
	   }

	   fmt.Println(obj.stack)

	   ok, el1 = obj.Top()
	   if ok {
	       fmt.Println("top = ", el1)
	   }
	*/

	obj2 := Constructor()
	fmt.Println(obj2.Size())
	fmt.Println(obj2.Empty())
	obj2.Push(10)
	obj2.Push(20)
	obj2.Push(30)
	obj2.Push(40)
	obj2.Print()
	fmt.Println("len =", obj2.Size())
	fmt.Println(obj2.Empty())
	el := obj2.Peek()
	fmt.Println("peek = ", el)
	el = obj2.Pop()
	fmt.Println("pop = ", el)
	obj2.Print()
	el = obj2.Pop()
	fmt.Println("pop = ", el)
	obj2.Print()

	os.Exit(1)
}
