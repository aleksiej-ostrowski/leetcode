package main

import (
	"fmt"
	// "unicode/utf8"
)

// https://www.educative.io/answers/how-to-implement-a-stack-in-golang

type Stack []int32

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(el int32) {
	*s = append(*s, el)
}

func (s *Stack) Pop() (int32, bool) {
	s_len := s.Len()

	if s_len == 0 {
		return 0, false
	}

	index := s_len - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) Last() (int32, bool) {
	s_len := s.Len()
	if s_len == 0 {
		return 0, false
	}
	index := s_len - 1
	element := (*s)[index]
	return element, true
}

func isValid(s string) bool {

	var stack Stack

	// fmt.Println(s)

	for _, el := range s {

		el_from_stack, check := stack.Last()

		good := false

		if check {
			if (int32('(') == el_from_stack) && (int32(')') == el) {
				_, _ = stack.Pop()
				good = true
			} else {
				if (int32('[') == el_from_stack) && (int32(']') == el) {
					_, _ = stack.Pop()
					good = true
				} else {
					if (int32('{') == el_from_stack) && (int32('}') == el) {
						_, _ = stack.Pop()
						good = true
					}
				}
			}
		}

		if !good {
			stack.Push(el)
		}

		// fmt.Printf("%d - %c \n", i, el)
	}

	return stack.Len() == 0
}

func main() {
	fmt.Println(isValid("{{{{{}}}}}}"))
}
