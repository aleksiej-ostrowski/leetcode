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
	"math"
)

func constructRectangle(area int) []int {
	min_1_d_2 := area
	length := area
	width := 1
	cand1 := 0
	margin := int(math.Ceil(math.Sqrt(float64(area))))
	for {
		cand1 += 1
		if cand1 > margin {
			break
		}
		if area%cand1 != 0 {
			continue
		}
		cand2 := area / cand1
		if cand2 >= cand1 {
			if (cand2 - cand1) < min_1_d_2 {
				length = cand2
				width = cand1
			}
		}
	}
	return []int{length, width}
}

func main() {
	fmt.Println(constructRectangle(1))
	// fmt.Println(constructRectangle(122122))
	fmt.Println("ok")
}
