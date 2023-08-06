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

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type TData struct {
	idx int
	cnt int
}

func findRestaurant(list1 []string, list2 []string) []string {
	rsl := make(map[string]TData)
	for idx, el := range list1 {
		rsl[el] = TData{idx: idx, cnt: rsl[el].cnt + 1}
	}
	mx := MaxInt
	var res []string
	for idx, el := range list2 {
		if rsl[el].cnt > 0 {
			sm := idx + rsl[el].idx
			if sm < mx {
				res = []string{el}
				mx = sm
			} else if sm == mx {
				res = append(res, el)
			}
		}
	}
	return res
}

func main() {
	list1 := []string{"happy", "sad", "good"}
	list2 := []string{"sad", "happy", "good"}
	fmt.Println(findRestaurant(list1, list2))
	fmt.Println("ok")
}
