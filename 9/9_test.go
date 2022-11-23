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
	arg      int
	expected bool
}

var allTests = []isTest{
	{1, true},
	{-1, false},
	{11, true},
	{121, true},
	{122, false},
	{1230321, true},
	{1230123, false},
	{11223344944332212, false},
	{11223344944332211, true},
	{-11223344944332211, false},
}

func TestisPalidrome(t *testing.T) {

	for _, test := range allTests {
		if output := isPalindrome(test.arg); output != test.expected {
			t.Errorf("Broke in %d", test.arg)
		}
	}
}
