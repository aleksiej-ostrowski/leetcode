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

func judgeCircle(moves string) bool {
	v := 0
	g := 0
	for _, el := range moves {
		switch el {
		case 82 /*'R'*/ :
			g += 1
		case 76 /*'L'*/ :
			g -= 1
		case 85 /*'U'*/ :
			v += 1
		case 68 /*'D'*/ :
			v -= 1
		}
		// fmt.Println(string(el))
	}
	return v == 0 && g == 0
}

func main() {
	fmt.Println(judgeCircle("UD"))
	fmt.Println("ok")
}
