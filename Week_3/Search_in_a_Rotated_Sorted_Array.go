package main

func search(nums []int, target int) int {

	// return -1 for an empty array
	if len(nums) == 0 {
		return -1
	}

	// get the pivot index
	// pivot is where the order of the array is lost
	start := 0
	end := len(nums) - 1
	var midpoint int

	for start != end {
		midpoint = start + (end-start)/2

		// if the section is still in order
		if nums[midpoint] > nums[start] {
			start++ // shift the midpoint by shifting the start
		} else {
			// focus between  the start to the midponit
			end = midpoint
		}
	}

	// check if the target is in the first section
	if target >= nums[0] && target <= nums[midpoint] {
		start = 0
		end = midpoint
	} else {
		// second section
		start = midpoint + 1
		end = len(nums) - 1
	}

	// perfrom binary search
	for start <= end {

		midpoint = start + (end-start)/2

		// return if found
		if nums[midpoint] == target {
			return midpoint
		} else {

			if target > nums[midpoint] {
				start = midpoint + 1
			} else {
				// focus between  the start to the midponit
				end = midpoint - 1
			}
		}
	}

	return -1

}

func main() {

	var i = []int{1, 3}
	println(search(i, 1))

}
