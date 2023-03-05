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

func isSubsequence(s string, t string) bool {
	len_s := len(s)
	len_t := len(t)

	if len_s == 0 {
		return true
	}

	if len_t == 0 {
		return false
	}

	cnt := 0
	idx_s := 0
	idx_t := 0

	for idx_s < len_s {
		for idx_t < len_t {
			if s[idx_s] == t[idx_t] {
				cnt += 1
				idx_t++
				break
			}
			idx_t++
		}
		idx_s++
	}

	return cnt == len_s
}

func main() {
	/*
	   s := "abc"
	   t := "ahbgdc"
	*/
	s := "aaaaaa"
	t := "bbaaaa"
	res := isSubsequence(s, t)
	fmt.Println(res)
	fmt.Println("ok")
}
