/*

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

*/

package main 

import (
    "testing"
)

type isTest struct {
    arg int
    expected bool
}

var allTests = []isTest{
    isTest{1, true},
    isTest{-1, false},
    isTest{11, true},
    isTest{121, true},
    isTest{122, false},
    isTest{1230321, true},
    isTest{1230123, false},
    isTest{11223344944332212, false},
    isTest{11223344944332211, true},
    isTest{-11223344944332211, false},
}

func TestisPalidrome(t *testing.T) {

    for _, test := range allTests {
        if output := isPalindrome(test.arg); output != test.expected {
            t.Errorf("Broke in %d", test.arg)
        }
    }
}
