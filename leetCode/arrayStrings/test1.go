package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	res := pivotIndex2(nums)
	fmt.Printf("res: %v\n", res)
}

func pivotIndex2(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}

	leftSum := 0
	for i, num := range arr {
		if leftSum == total-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
