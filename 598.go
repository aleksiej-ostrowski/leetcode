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

const (
    intSize = 32 << (^uint(0) >> 63)
    MaxInt = 1 << (intSize - 1) - 1
)

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxCount(m int, n int, ops [][]int) int {
    if len(ops) == 0 {
        return m * n
    }
    min_x := MaxInt
    min_y := MaxInt
    for _, el := range ops {
        min_x = Min(min_x, el[0])
        min_y = Min(min_y, el[1])
    }
    return min_x * min_y
}

func main() {
    fmt.Println(maxCount(3, 3, [][]int{{2,2},{3,3}}))
	fmt.Println("ok")
}
