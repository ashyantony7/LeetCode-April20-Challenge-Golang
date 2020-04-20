package main

import "fmt"

func isHappy(n int) bool {

	// map to see if the element occurance
	var set = make(map[int]int)
	set[n] = 1 // add the first element

	for n != 1 {
		var temp int = n // make a temp copy

		var sum int
		for temp != 0 {
			sum += (temp % 10) * (temp % 10)
			temp /= 10
		}

		// if result already exists quit
		if _, found := set[sum]; found {
			return false
		} else {
			set[sum] = 1
		}

		n = sum // update n

	}
	return true
}

func main() {

	i := 19
	fmt.Printf("%v\n", isHappy(i))

}
