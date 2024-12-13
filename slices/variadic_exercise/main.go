package main

func sum(nums ...int) int {
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}

	return totalSum
}
