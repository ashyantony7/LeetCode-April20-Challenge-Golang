package main

func countElements(arr []int) int {

	// map for frequency
	var count = make(map[int]int)

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	// loop again to count (x-1) elements
	var res int
	for i, _ := range count {
		res += count[i-1]
	}
	return res
}
