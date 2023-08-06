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
	// "strconv"
	// "math"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getBits32(val int32) [32]uint8 {
	// fmt.Println(strconv.FormatInt(val, 2))
	var res [32]uint8
	for bit := 0; bit < 32; bit++ {
		res[31-bit] = uint8((val >> bit) & 1)
	}
	return res
}

func getHamming(a, b [32]uint8) int {
	res := 0
	for bit := 0; bit < 32; bit++ {
		res += Abs(int(a[bit]) - int(b[bit]))
	}
	return res
}

func hammingDistance(x int, y int) int {
	// int32  : -2147483648 to 2147483647
	return getHamming(getBits32(int32(x)), getBits32(int32(y)))
}

func main() {
	//fmt.Println(getBits32(12345))
	//fmt.Println(strconv.FormatInt(12345, 2))
	fmt.Println(hammingDistance(3, 1))
	fmt.Println("ok")
}
