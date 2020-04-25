package main

import "fmt"

type stack []rune

// IsEmpty: check if stack is empty
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *stack) Push(n rune) {
	*s = append(*s, n)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func backspaceCompare(S string, T string) bool {

	var sStack, tStack stack

	// push elements into stack
	// pop them if backspaces occur
	for _, s := range S {
		if s != rune('#') {
			sStack.Push(s)
		} else {
			_, _ = sStack.Pop()
		}
	}

	for _, t := range T {
		if t != rune('#') {
			tStack.Push(t)
		} else {
			_, _ = tStack.Pop()
		}
	}

	if len(sStack) != len(tStack) {
		return false
	} else if len(sStack) == 0 {
		return true
	} else {

		for !sStack.IsEmpty() {
			s, _ := sStack.Pop()
			t, _ := tStack.Pop()

			if s != t {
				return false
			}
		}

		return true
	}
}

func main() {

	s := "abcd"
	t := "bbcd"

	fmt.Println(backspaceCompare(s, t))
}
