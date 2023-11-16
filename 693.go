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

func hasAlternatingBits(n int) bool {
	first := true
	pred := uint8(0)
	for {
		if !first {
			n >>= 1
		}
		if n == 0 {
			break
		}
		new_bit := uint8(n & 1)
		// fmt.Println(new_bit)
		if first {
			pred = new_bit
			first = false
		} else {
			if pred == new_bit {
				return false
			}
			pred = new_bit
		}
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
}
