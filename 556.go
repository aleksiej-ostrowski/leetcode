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

func matrixReshape(mat [][]int, r int, c int) [][]int {
    if len(mat) * len(mat[0]) != r * c {
        return mat 
    }
    ch := make(chan int, r * c)
    for col:= 0; col < len(mat); col += 1 {
        for row:= 0; row < len(mat[col]); row += 1 {
            ch<- mat[col][row]    
        }
    }
    res := make([][]int, r)
    for row:= 0; row < r; row += 1 {
        res[row] = make([]int, c)
        for col:= 0; col < c; col += 1 {
            res[row][col] = <-ch    
        }
    }
    close(ch)
    return res
}

func main() {
    mat := [][]int{{1,2},{3,4},}
    fmt.Println(mat)
    res := matrixReshape(mat, 1, 4)
    fmt.Println(res)
	fmt.Println("ok")
}
