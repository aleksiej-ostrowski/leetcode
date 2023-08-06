// Развернуть односвязный список

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

type TLinkNode struct {
	next  *TLinkNode
	value int
}

func prepareLinkNode(data []int) *TLinkNode {
	node := &TLinkNode{nil, 0}
	node_start := node
	for idx, el := range data {
		node.value = el
		if idx != len(data)-1 {
			node.next = &TLinkNode{nil, 0}
			node = node.next
		}
	}
	return node_start
}

func printLinkNode(node *TLinkNode) {
	new_node := node
	for {
		if new_node == nil {
			break
		}
		fmt.Println(new_node.value)
		new_node = new_node.next
	}
}

func invertLinkNode(node *TLinkNode) *TLinkNode {
	new_node := node
	res_node := &TLinkNode{nil, 0}
	cnt := 0
	for {
		if new_node == nil {
			break
		}
		if cnt != 0 {
			res_node = &TLinkNode{res_node, new_node.value}
		} else {
			res_node = &TLinkNode{nil, new_node.value}
		}
		cnt += 1
		new_node = new_node.next
	}
	return res_node
}

func main() {
	new_node := prepareLinkNode([]int{3, 10, 34, 51, 101, 202, 150})
	fmt.Println("Data: ")
	printLinkNode(new_node)
	res_node := invertLinkNode(new_node)
	fmt.Println("Result: ")
	printLinkNode(res_node)
	fmt.Println("ok")
}
