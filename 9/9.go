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

    nums := []int8{}

    for (x > 0) {
        nums = append(nums, int8(x % 10))
        x /= 10
    }

    sum := 0
    for i, j := 0, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
        // fmt.Println("nums[", i, "] = ", nums[i], " nums[", j, "] = ", nums[j])
        if nums[i] == nums[j] {
            sum += 1
        }
    }

    // fmt.Println("len(nums) = ", len(nums), " sum = ", sum)

    return (len(nums) >> 1) == sum
}

func main() {
	fmt.Println(isPalindrome(123321))
}
