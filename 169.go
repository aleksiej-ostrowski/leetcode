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

func majorityElement(nums []int) int {
	len_nums := len(nums)
	cnt := len_nums >> 1

	memory := make(map[int]int)
	for _, el := range nums {
		memory[el]++
	}

	for key, val := range memory {
		if val > cnt {
			return key
		}
	}

	return 0
}

func main() {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(arr))
}
