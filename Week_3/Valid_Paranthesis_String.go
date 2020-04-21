package main

import (
	"fmt"
)

type stack []int

// IsEmpty: check if stack is empty
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *stack) Push(n int) {
	*s = append(*s, n)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func checkValidString(s string) bool {

	var left stack
	var asterisk stack

	for i, val := range s {

		if val == '(' {
			left.Push(i)
		} else if val == ')' {
			if !left.IsEmpty() {
				_, _ = left.Pop()
			} else if !asterisk.IsEmpty() {
				_, _ = asterisk.Pop()
			} else {
				return false
			}
		} else if val == '*' {
			asterisk.Push(i)
		}
	}

	for !left.IsEmpty() {
		if !asterisk.IsEmpty() {

			i, _ := asterisk.Pop()
			l, _ := left.Pop()

			if l > i {
				return false
			}

		} else {
			return false
		}
	}

	return true
}

func main() {

	str := "((()))*"

	fmt.Printf("%v ", checkValidString(str))
}
