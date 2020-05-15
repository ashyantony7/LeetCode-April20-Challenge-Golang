package main

type MinStack struct {
	stack  []int
	minVal []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := new(MinStack)
	m.stack = make([]int, 1)
	m.minVal = make([]int, 1)
	return *m
}

func (this *MinStack) Push(x int) {

	// for an empty stack
	if len(this.stack) == 0 {

		// add values to both stacks
		this.stack = append(this.stack, x)
		this.minVal = append(this.minVal, x)
	} else {

		// add to value stack
		this.stack = append(this.stack, x)

		// update min value stack
		last := len(this.minVal) - 1
		if x < this.minVal[last] {
			this.minVal = append(this.minVal, x)
		} else {
			this.minVal = append(this.minVal, this.minVal[last])
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) != 0 {
		index := len(this.stack) - 1      // Get the index of the top most element.
		this.stack = (this.stack)[:index] // Remove values from both stacks by slicing it off.
		this.minVal = (this.minVal)[:index]
	}
}

func (this *MinStack) Top() int {

	last := len(this.stack) - 1
	return this.stack[last]
}

func (this *MinStack) GetMin() int {

	last := len(this.minVal) - 1
	return this.minVal[last]
}

func main() {

	s := new(MinStack)
	s.Push(2)
	s.Push(0)
	s.Push(3)
	s.Push(0)
	println(s.GetMin())
	s.Pop()
	println(s.GetMin())
	s.Pop()
	println(s.GetMin())
	s.Pop()
	println(s.GetMin())
}
