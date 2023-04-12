package main

import "fmt"

func main() {
	nums := []int{-4, 3, -9, 0, 4, 1}
	positive := []int{}
	negative := []int{}
	zeros := []int{}
	for _, num := range nums {
		if num > 0 {
			positive = append(positive, num)
		} else if num < 0 {
			negative = append(negative, num)
		} else {
			zeros = append(zeros, num)
		}
	}
	n := float32(len(nums))
	var minipositive = float32(len(positive)) / n
	var mininegative = float32(len(negative)) / n
	var minizeros = float32(len(zeros)) / n

	fmt.Printf("%.6f \n%.6f \n%.6f", minipositive, mininegative, minizeros)
}
