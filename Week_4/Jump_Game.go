package main

func canJump(nums []int) bool {

	// start with last element
	var pos int = len(nums) - 1

	// go from end to start and see if the jump
	// could have been made
	for i := (len(nums) - 2); i >= 0; i-- {
		if nums[i] >= (pos - i) {
			pos = i
		}
	}

	if pos == 0 {
		return true
	} else {
		return false
	}

}
