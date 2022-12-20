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

func singleNumber(nums []int) int {
	memory := make(map[int]uint8)

	for _, el := range nums {
		memory[el]++
	}

	for key, val := range memory {
		if val == 1 {
			return key
		}
	}

	return 0
}

func main() {
	arr := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(arr))
}
