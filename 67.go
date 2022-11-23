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

func charbin2byte(char byte) byte {
	switch char {
	case 48:
		return 0
	case 49:
		return 1
	}
	return 255
}

func byte2charbin(b byte) byte {
	switch b {
	case 0:
		return 48
	case 1:
		return 49
	}
	return 255
}

func addBinary(a string, b string) string {
	overflow := 0

	len_a := len(a)
	len_b := len(b)
	ind_a := len_a - 1
	ind_b := len_b - 1

	var res []rune

	if len_a > len_b {
		res = make([]rune, len_a)
	} else {
		res = make([]rune, len_b)
	}

	len_res := len(res)
	ind_res := len_res - 1

	for {

		new_var := overflow

		if ind_a >= 0 {
			new_var += int(charbin2byte(a[ind_a]))
			ind_a -= 1
		}

		if ind_b >= 0 {
			new_var += int(charbin2byte(b[ind_b]))
			ind_b -= 1
		}

		overflow = 0

		switch new_var {
		case 0, 1:
			res[ind_res] = rune(byte2charbin(byte(new_var)))
		case 2:
			res[ind_res] = rune(byte2charbin(0))
			overflow += 1
		case 3:
			res[ind_res] = rune(byte2charbin(1))
			overflow += 1
		}

		if (ind_a < 0) && (ind_b < 0) {
			break
		}

		ind_res -= 1
	}

	if overflow > 0 {
		res = append([]rune{49}, res...)
	}

	return string(res)
}

func main() {
	a := "10111"
	b := "1"

	fmt.Println(addBinary(a, b))
	fmt.Println("ok")
}
