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

func Min(a, b int) int {
    if a < b {
        return a
    } 
    return b
}

func distributeCandies(candyType []int) int {    
    memory := make(map[int]bool)    
    for _, el := range candyType {    
        memory[el] = true    
    }    
    return Min(len(memory), len(candyType) >> 1)    
}

func main() {
    fmt.Println(distributeCandies([]int{1,1,2,2,3,3}))
	fmt.Println("ok")
}
