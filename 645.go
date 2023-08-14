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

func findErrorNums(nums []int) []int {
	al := make(map[int]bool)
	for ind := 1; ind <= len(nums); ind++ {
		al[ind] = true
	}
	res := 0
	for _, el := range nums {
		if _, ok := al[el]; ok {
			delete(al, el)
		} else {
			res = el
		}
	}
	// fmt.Println(al)
	for key := range al {
		return []int{res, key}
	}
	return []int{0, 0}
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println("ok")
}
