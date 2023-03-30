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

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	memory := make(map[int]int)
	for _, el := range nums1 {
		memory[el] = -1
	}
	idx := -1
	len_nums2 := len(nums2)
	for {
		idx += 1
		if idx >= len_nums2 {
			break
		}
		el := nums2[idx]
		if val, ok := memory[el]; ok {
			if val == -1 {
				memory[el] = el
				continue
			}
		}

		for key, val := range memory {
			if val == -1 {
				continue
			}
			if el > val {
				memory[key] = el
			}
		}
	}
	var res []int
	for _, el := range nums1 {
		if val, ok := memory[el]; ok {
			if val != el {
				res = append(res, val)
			} else {
				res = append(res, -1)
			}
		}
	}

	return res
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println("ok")
}
