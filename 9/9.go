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
	// "math"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	var new_x int

	x_ := x

	for x_ > 0 {
		a := x_ / 10
		new_x = (new_x<<3 + new_x<<1) + (x_ - (a<<3 + a<<1))
		x_ = a
	}

	return new_x == x
}

func main() {
	fmt.Println(isPalindrome(123321))
}
