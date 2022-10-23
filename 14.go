/*

#--------------------------------#
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
#                                #
#  https://aleksiej.com          #
#                                #
#--------------------------------#  

*/

package main

import (
    "fmt"
)

func longestCommonPrefix(strs []string) string {
    if (len(strs) == 0) {
        return ""
    }
    maxlen := 201
    exclude := 0
    for i, el := range strs {
        len_el := len(el)
        if (len_el == 0) {
            return ""
        }
        if (len_el < maxlen) {
            exclude = i
            maxlen = len_el
        }
    }

    min_ := maxlen
    first := &strs[exclude]
    for i, el := range strs {
        if (i == exclude) {
            continue
        }    
        sum := 0
        for j:=0; j < maxlen; j++ {
            if ((*first)[j] == el[j]) {
                sum += 1
            } else {
                break     
            }
        }
        if (sum == 0) {
            return ""
        }

        if (sum < min_) {
           min_ = sum 
        }
    }    
    return (*first)[:min_]
}

func main() {
    var strs = []string{"f"}
    fmt.Println(longestCommonPrefix(strs))
}



