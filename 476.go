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

func binToInt32(arr [32]uint8) int32 {
    var res int32
    // fmt.Println(arr)
    for idx:=0;idx<32;idx++ {
        res += int32(arr[31-idx]) * (1 << idx)
    }
    return res 
}


func binComplement(arr [32]uint8) [32]uint8 {
    start := 0
    //fmt.Println(arr)
    for {
        if arr[start] == 0 {
            if start + 1 < 32 {
                start += 1
            } else {
                break
            }
        } else {
            break
        }
    }
    for idx:=start;idx<32;idx++ {
        arr[idx] ^= 1
    }
    //fmt.Println(arr)
    return arr
}


func int32ToBin(num int32) [32]uint8 {
    var res [32]uint8
    for idx:=0;idx<32;idx++ {
        res[31 - idx] = uint8((num >> idx) & 1)
    }
    return res
}

func findComplement(num int) int {
// int32  : -2147483648 to 2147483647
    return int(binToInt32(binComplement(int32ToBin(int32(num)))))
}

func main() {
    // fmt.Println(int32ToBin(5))
    fmt.Println("result = ", findComplement(5))
	fmt.Println("ok")
}
