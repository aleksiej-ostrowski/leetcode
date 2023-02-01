
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

func isPowerOfTwo(n int) bool {
    if n <= 0 {
        return false
    }    
    var eps float64 = 1000000000.0
    var num float64 = math.Log2(float64(n))
    return int(math.Trunc(num * eps)) == int(math.Trunc(num) * eps)
}

func main() {
    fmt.Println(isPowerOfTwo(17))
	fmt.Println("ok")
}
