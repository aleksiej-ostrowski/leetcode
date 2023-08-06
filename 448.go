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

func findDisappearedNumbers(nums []int) []int {
	memory := make(map[uint32]bool)
	max_el := len(nums)
	for _, el := range nums {
		memory[uint32(el)] = true
	}
	res := []int{}
	for el := 1; el <= max_el; el++ {
		if _, ok := memory[uint32(el)]; !ok {
			res = append(res, el)
		}
	}
	return res
}

func main() {
	res := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(res)
}
