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
    "sync"
    "sync/atomic"
    "unicode"
)

func IsUpper(s rune) bool {
    return unicode.IsUpper(s) && unicode.IsLetter(s)
}

func IsLower(s rune) bool {
    return unicode.IsLower(s) && unicode.IsLetter(s)
}

func detectCapitalUse(word string) bool {
    first := int64(0)
    all_lower := int64(0)
    all_upper := int64(0)

    var wg sync.WaitGroup

    for idx, el := range word {
        wg.Add(1)
        go func(idx int, el rune) {
            defer wg.Done()
            up := IsUpper(el)
            if up {
                atomic.AddInt64(&all_upper, 1)
            }
            if IsLower(el) {
                atomic.AddInt64(&all_lower, 1)
            }
            if idx == 0 && up {
                atomic.AddInt64(&first, 1)
            }
        }(idx, el)
   }
   wg.Wait()
   len_word := int64(len(word))
   return (first == 1 && all_lower == len_word-1) || (all_upper == len_word) || (all_lower == len_word)   
}

func main() {
    fmt.Println(detectCapitalUse("USA"))
	fmt.Println("ok")
}


