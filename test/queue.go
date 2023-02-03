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

type MyQueue struct {
	queue []int
}

func Constructor() MyQueue {
	queue := make([]int, 0)
	return MyQueue{queue}
}

func (this *MyQueue) PushToBack(x int) {
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

func main() {
	obj := Constructor()
	fmt.Println(obj.Size())
	fmt.Println(obj.Empty())
	obj.PushToBack(10)
	obj.PushToBack(20)
	obj.PushToBack(30)
	obj.PushToBack(40)
	fmt.Println(obj.queue)
	fmt.Println(obj.Size())
	fmt.Println(obj.Empty())
	ok, el := obj.PeekFromFront()
	fmt.Println(ok, el)
	ok, el = obj.PopFromFront()
	fmt.Println(ok, el)
	fmt.Println(obj.queue)
}
