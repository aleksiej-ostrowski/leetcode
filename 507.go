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

func checkPerfectNumber(num int) bool {
    sum := 1
    for nm:= 2; nm <= num >> 1; nm++ {
        if num % nm == 0 {
            sum += nm
        }
    }
    return sum == num
}

func main() {
    fmt.Println(checkPerfectNumber(7))
    // fmt.Println(checkPerfectNumber(28))
	fmt.Println("ok")
}
