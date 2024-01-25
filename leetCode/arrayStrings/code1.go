package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	res := pivotIndex(nums)
	fmt.Printf("res: %v\n", res)
}

// todo: 时间复杂度为 O(n)，空间复杂度为 O(1)
func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == total-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
