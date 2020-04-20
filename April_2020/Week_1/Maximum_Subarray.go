package main

func maxSubArray(nums []int) int {

	var maxVal int = nums[0]
	var sum int = nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] > (nums[i] + sum) {
			if nums[i] > maxVal {
				maxVal = nums[i]
			}
			sum = nums[i]
		} else {
			if (nums[i] + sum) > maxVal {
				maxVal = nums[i] + sum
			}
			sum += nums[i]
		}
	}
	return maxVal
}
