package main

import (
	"fmt"
	"sort"
)

// sort
func sum(nums []int32) int64 {
	var result int64
	for _, num := range nums {
		result += int64(num)
	}
	return result
}
func solve(nums []int32) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println()
	fmt.Printf("%d %d", sum(nums[0:4]), sum(nums[1:5]))

}
func main() {
	nums := []int32{1, 2, 3, 4, 5}
	solve(nums)
}
