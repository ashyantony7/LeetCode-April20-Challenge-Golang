package main

func moveZeroes(nums []int) {

	// keep an index for non zero elements
	var index int = 0

	for i := 0; i < len(nums); i++ {

		// if non zero just copy front
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}

	// fill the rest with zeros
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
