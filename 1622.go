/*

#-------------------------------------------------------------#
#                                                             #
#  version 0.0.2                                              #
#  https://leetcode.com/problems/fancy-sequence/              #
#                                                             #
#  Aleksiej Ostrowski, 2021                                   #
#                                                             #
#  https://aleksiej.com                                       #
#                                                             #
#-------------------------------------------------------------#

*/

package main

import (
	"fmt"
	"os"
)

const crt = 84467440737095516 // max(uint64) = 8446744073709551615
const md = 1000000007

type Oper struct {
	o int8
	v int8
	h int32
}

type Fancy struct {
	arr []int8
	opr []Oper
	cas map[int]int
}

func Constructor() Fancy {
	return Fancy{}
}

func (this *Fancy) Append(val int) {
	this.arr = append(this.arr, int8(val))
}

func printList(l Fancy) {

	for _, v := range l.arr {

		fmt.Print(v, " ")
	}

	fmt.Println()
}

func (this *Fancy) AddAll(inc int) {

	this.opr = append(this.opr, Oper{-2, int8(inc), int32(len(this.arr))})
	this.cas = make(map[int]int)

	/*

	   for k := range this.cas {
	       delete(this.cas, k)
	   }

	*/
}

func (this *Fancy) MultAll(m int) {

	this.opr = append(this.opr, Oper{-1, int8(m), int32(len(this.arr))})
	this.cas = make(map[int]int)

	/*

	   for k := range this.cas {
	       delete(this.cas, k)
	   }

	*/
}

func (this *Fancy) GetIndex(idx int) int {

	if idx < 0 || idx >= len(this.arr) {
		return -1
	}

	if val, ok := this.cas[idx]; ok {
		return val
	}

	var vv uint64

	vv = uint64(this.arr[idx])

	for _, v := range this.opr {

		if idx >= int(v.h) {
			continue
		}

		switch os := v.o; os {
		case -2:
			vv += uint64(v.v)
		case -1:
			vv *= uint64(v.v)
		}

		if vv > crt {
			vv = vv % md
		}
	}

	if vv > md {
		vv = vv % md
	}

	if this.cas == nil {
		this.cas = make(map[int]int)
	}

	this.cas[idx] = int(vv)

	return int(vv)
}

func main() {

	obj := Constructor()

	obj.Append(2)
	obj.Append(4)
	obj.Append(1)
	obj.Append(6)

	printList(obj)

	obj.AddAll(9)
	obj.AddAll(1)
	obj.MultAll(2)

	fmt.Println(obj.GetIndex(0))

	os.Exit(1)

	obj.MultAll(9)
	obj.AddAll(2)
	fmt.Println(obj.GetIndex(3))
	obj.Append(7)
	obj.Append(5)
	obj.Append(3)
	obj.AddAll(4)
	obj.MultAll(5)
	obj.AddAll(4)
	fmt.Println(obj.GetIndex(1))
}
