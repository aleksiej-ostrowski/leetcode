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

func removeDuplicates(nums []int) int {
	cur := 1
	len_nums := len(nums)
	// fmt.Println(nums)
	i := 0
	for i < len_nums {
		j := i + 1
		for ; j < len_nums; j++ {
			if nums[i] != nums[j] {
				break
			}
		}
		if j < len_nums {
			nums[cur] = nums[j]
			cur += 1
			i = j - 1
		}
		i += 1
	}
	// for j := cur; j < len_nums; j++ {
	// 	nums[j] = -1
	// }
	// fmt.Println(nums)
	return cur
}

func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
