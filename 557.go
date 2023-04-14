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
    "strings"
)

func reverseWords(s string) string {
    res := []string{}
    buf := ""

    for _, el := range s {
        if el != ' ' {
            buf = string(el) + buf
        } else {
            res = append(res, buf)
            buf = ""
        }
    }

    if buf != "" {
        res = append(res, buf)
    }

    return strings.Join(res, " ")
}


func main() {
    fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println("ok")
}
