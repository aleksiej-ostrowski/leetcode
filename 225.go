
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

type MyQueue struct {
    queue []int
}

func ConstructorQ() MyQueue {
    queue := make([]int, 0)
    return MyQueue{queue}
}

func (this *MyQueue) PushToBack(x int)  {
    this.queue = append(this.queue, x)
}

func (this *MyQueue) PeekFromFront() (bool, int) {
    if this.Empty() {
        return false, 0
    }    
    return true, this.queue[0]
}

func (this *MyQueue) PopFromFront() (bool, int) {
    if this.Empty() {
        return false, 0
    }    
    el := this.queue[0]
    this.queue = this.queue[1:]
    return true, el
}

func (this *MyQueue) Size() int {
    return len(this.queue)
}

func (this *MyQueue) Empty() bool {
    return this.Size() == 0
}



type MyStack struct {
    q1 MyQueue
    q2 MyQueue
    current int 
}


func Constructor() MyStack {
    var stack MyStack 
    stack.q1 = ConstructorQ()
    stack.q2 = ConstructorQ()
    stack.current = 1
    return stack
}


func (this *MyStack) Push(x int)  {
    switch this.current {
    case 1:
        if this.q1.Size() > 0 {
            this.q2.PushToBack(x)
            for {
                ok, el := this.q1.PopFromFront()
                if ok {
                    this.q2.PushToBack(el)
                } else {
                    break
                }
            }
            this.current = 2
        } else {
            this.q1.PushToBack(x)
        }
    case 2:
        if this.q2.Size() > 0 {
            this.q1.PushToBack(x)
            for {
                ok, el := this.q2.PopFromFront()
                if ok {
                    this.q1.PushToBack(el)
                } else {
                    break
                }
            }
            this.current = 1
        } else {
            this.q2.PushToBack(x)
        }
    }    
}

func (this *MyStack) Pop() int {
    switch this.current {
    case 1:
        _, el := this.q1.PopFromFront()
        if this.q1.Size() == 0 {
            this.current = 2
        }
        return el
    case 2:
        _, el := this.q2.PopFromFront()
        if this.q2.Size() == 0 {
            this.current = 1
        }
        return el
    }   
    return -1
}


func (this *MyStack) Top() int {
    switch this.current {
    case 1:
        _, el := this.q1.PeekFromFront()
        return el
    case 2:
        _, el := this.q2.PeekFromFront()
        return el
    }    
    return -1
}


func (this *MyStack) Empty() bool {
    return this.q1.Size() + this.q2.Size() == 0
}

func (this *MyStack) Print()  {
    memory := MyStack{this.q1, this.q2, this.current}
    switch memory.current {
    case 1:
        if memory.q1.Size() > 0 {
            for {
                ok, el := memory.q1.PopFromFront()
                if ok {
                    fmt.Println("el = ", el)
                } else {
                    break
                }
            }
        }
    case 2:
        if memory.q2.Size() > 0 {
            for {
                ok, el := memory.q2.PopFromFront()
                if ok {
                    fmt.Println("el = ", el)
                } else {
                    break
                }
            }
        }
     } 
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */


func main() {
    /*
    obj := Constructor()
    obj.Push(10)
    p2 := obj.Pop()
    p3 := obj.Top()
    p4 := obj.Empty()
    fmt.Println(p2, p3, p4)
    fmt.Println("ok")
    */
    obj := Constructor()
    obj.Push(10)
    obj.Push(20)
    obj.Push(30)
    obj.Print()

    el1 := obj.Pop()
    fmt.Println("pop = ", el1)

    obj.Print()
    
    obj.Push(40)

    el1 = obj.Pop()
    fmt.Println("pop = ", el1)

    obj.Print()
    os.Exit(1)

    fmt.Println("top = ", obj.Top())
    fmt.Println(obj.Empty())
    obj.Print()
    os.Exit(1)

    obj2 := ConstructorQ()
    fmt.Println(obj2.Size())
    fmt.Println(obj2.Empty())
    obj2.PushToBack(10)
    obj2.PushToBack(20)
    obj2.PushToBack(30)
    obj2.PushToBack(40)
    fmt.Println(obj2.queue)
    fmt.Println(obj2.Size())
    fmt.Println(obj2.Empty())
    ok, el := obj2.PeekFromFront()
    fmt.Println(ok, el)
    ok, el = obj2.PopFromFront()
    fmt.Println(ok, el)
    fmt.Println(obj2.queue)
 }
