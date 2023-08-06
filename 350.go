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

func intersect(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)

	for _, el1 := range nums1 {
		map1[el1] += 1
	}

	res := []int{}
	for _, el2 := range nums2 {
		if val, ok := map1[el2]; ok {
			if val <= 0 {
				delete(map1, el2)
			} else {
				res = append(res, el2)
				map1[el2] -= 1
			}
		}
	}

	return res
}

func main() {
	// a := []int{1,2,2,2,1}
	// b := []int{2,2}
	a := []int{4, 9, 5}
	b := []int{9, 4, 9, 8, 4}
	res := intersect(a, b)
	fmt.Println(res)
	fmt.Println("ok")
}
