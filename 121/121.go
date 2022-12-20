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

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	len_prices := len(prices)
	ind0 := 0
	ind1 := len_prices - 1
	min := prices[ind0]
	max := prices[ind1]
	dd := max - min
	for ind := 1; ind < len_prices-1; ind++ {
		ind0++
		ind1--
		min_ := prices[ind0]
		max_ := prices[ind1]

		el := prices[ind]
		if el < min {
			min_ := el
			ind0_ := ind
		}
		if el > max {
			max_ := el
			ind1_ := ind
		}
		dd_ := max

	}
	dd := max - min
	if dd <= 0 {
		return 0
	}
	return dd
}

func main() {
	prices := []int{7, 6, 4, 3, 1} // {7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
