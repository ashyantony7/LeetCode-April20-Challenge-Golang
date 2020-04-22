package main

import "fmt"

func subarraySum(nums []int, k int) int {

	var res, sum int
	// make a hasmap to store the sum at
	// different points
	var count = make(map[int]int)

	for i := 0; i < len(nums); i++ {

		// keep cumulative sum
		sum += nums[i]

		// if the current sum matches k
		if sum == k {
			res++
		}

		// check for sum -k occurance also
		// take care of 0
		if val, found := count[sum-k]; found {
			res += val
		}

		// add the sum to Hashmap
		count[sum]++
	}

	return res
}

func main() {

	var arr = []int{1, 1, 1}

	fmt.Printf("%v \n", subarraySum(arr, 2))
}
