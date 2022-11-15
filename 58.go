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

func lengthOfLastWord(s string) int {
	sum_lst := 0
	sum_sav := 0
	for _, el := range s {
		if el == ' ' {
			if sum_lst > 0 {
				sum_sav = sum_lst
				sum_lst = 0
			}
		} else {
			sum_lst += 1
		}
	}
	if sum_lst > 0 {
		return sum_lst
	}
	return sum_sav
}

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}
