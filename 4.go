/*

#-------------------------------------------------------------#
#                                                             #
#  version 0.0.1                                              #
#  https://leetcode.com/problems/median-of-two-sorted-arrays/ #
#                                                             #
#  Aleksiej Ostrowski, 2020                                   #
#                                                             #
#  https://aleksiej.com                                       #
#                                                             #
#-------------------------------------------------------------#

*/

package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	i1 := 0
	i2 := 0

	var x1 []int
	var x2 []int

	x1 = nums1
	x2 = nums2

	switch {
	case (len(nums1) == 0) && (len(nums2) > 0):
		x1 = nums2
		x2 = nums1
	case (len(nums1) > 0) && (len(nums2) > 0):
		if nums1[i1] > nums2[i2] {
			x1 = nums2
			x2 = nums1
		}
	}

	r := make([]int, len(x1)+len(x2))

	ir := 0

	for {

		if i1 >= len(x1) {
			break
		}

		e1 := x1[i1]

		r[ir] = e1

		ir += 1

		e3 := e1

		flag := true

		if i1+1 < len(x1) {
			e3 = x1[i1+1]
			flag = false
		}

		for {

			if i2 >= len(x2) {
				break
			}

			e2 := x2[i2]
			if (e2 >= e1) && ((e2 <= e3) || flag) {
				r[ir] = e2
				ir += 1
				i2 += 1
			} else {
				break
			}
		}

		i1 += 1
	}

	ifz := len(r)%2 == 0

	fmt.Println(r)

	ind := len(r) / 2

	fmt.Println("ind = ", ind, "el = ", r[ind])

	if ifz {
		return (float64(r[ind-1]) + float64(r[ind])) * 0.5
	} else {
		return float64(r[ind])
	}

}

func main() {

	a := []int{1, 34, 56, 100, 200, 201, 500}
	b := []int{250}

	// a := []int{1,2,3,4,5}
	// b := []int{}

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println(findMedianSortedArrays(a, b))

}
