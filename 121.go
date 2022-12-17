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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/1735550/python-javascript-easy-solution-with-very-clear-explanation/
func maxProfit(prices []int) int {
	left := 0
	right := 1
	max_profit := 0
	len_prices := len(prices)
	for {
		if right >= len_prices {
			break
		}
		currentProfit := prices[right] - prices[left]
		if prices[left] < prices[right] {
			max_profit = max(currentProfit, max_profit)
		} else {
			left = right
		}
		right += 1
	}
	return max_profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
