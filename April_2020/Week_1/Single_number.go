package main

func singleNumber(nums []int) int {

	var count int

	for i := 0; i < len(nums); i++ {
		count = count ^ nums[i]
	}

	return count
}
