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
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	rnk := append(make([]int, 0, len(score)), score...)
	sort.Ints(rnk)
	mp := make(map[int]string)
	for idx, el := range rnk {
		new_idx := len(rnk) - 1 - idx
		switch new_idx {
		case 0:
			mp[el] = "Gold Medal"
		case 1:
			mp[el] = "Silver Medal"
		case 2:
			mp[el] = "Bronze Medal"
		default:
			mp[el] = strconv.FormatInt(int64(new_idx+1), 10)
		}
	}
	var res []string
	for _, el := range score {
		res = append(res, mp[el])
	}
	return res
}

func main() {
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
	fmt.Println("ok")
}
