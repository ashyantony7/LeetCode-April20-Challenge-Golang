package main

func findMaxLength(nums []int) int {

	var sum, res int
	var m = make(map[int]int)

	// start with negative index in map
	m[0] = -1

	for i := 0; i < len(nums); i++ {

		// -1 for 0
		if nums[i] == 0 {
			sum -= 1
		} else {
			sum++
		}

		// if already there find length
		if val, found := m[sum]; found {
			if (i - val) > res {
				res = (i - val)
			}
		} else {
			m[sum] = i
		}

	}

	return res
}
