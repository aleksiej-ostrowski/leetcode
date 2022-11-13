package main

import (
    "fmt"
)

func removeElement(nums []int, val int) int {
    cur := 0
    // fmt.Println(nums)
    for i, el := range nums {
        if el != val {
            if i != cur {
                nums[cur] = el
            }
            cur += 1
        }    
    }
    nums = nums[:cur]
    // fmt.Println(nums)
    return cur
}

func main() {
    nums := []int{10,3,3,10}
    val := 3
    fmt.Println(removeElement(nums, val))
    fmt.Println("ok")
}
