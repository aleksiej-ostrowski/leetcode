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

func intersection(nums1 []int, nums2 []int) []int {
    map1 := make(map[int]int)

    for _, el1 := range nums1 {
        map1[el1] += 1
    }    

    res := []int{}
    for _, el2 := range nums2 {
        if _, ok := map1[el2]; ok {
            res = append(res, el2)
            delete(map1, el2)
        }    
    }
        
    return res
}

func main() {
    a := []int{1,2,3,4,6,7,10}
    b := []int{0,3,4,6,15}
    res := intersection(a, b)
    fmt.Println(res)
	fmt.Println("ok")
}
