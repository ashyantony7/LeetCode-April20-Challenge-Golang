package main

func productExceptSelf(nums []int) []int {

	var res = make([]int, len(nums))
	var left = make([]int, len(nums))
	var right int = 1

	// array to store the results of left elements
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = (left[i-1] * nums[i-1])
	}

	// for the last element only left element product
	res[len(nums)-1] = left[len(nums)-1]

	// multiply right element product with other elements
	for i := (len(nums) - 2); i >= 0; i-- {
		right = right * nums[i+1]
		res[i] = left[i] * right
	}

	return res

}
