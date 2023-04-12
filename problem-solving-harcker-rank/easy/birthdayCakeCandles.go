package main

import (
	"fmt"
	"sort"
)

func birthdayCakeCandles(candles []int32) int32 {
	var result int32 = 1
	// Write your code here
	sort.Slice(candles, func(i, j int) bool {
		return candles[i] > candles[j]
	})
	var largest = candles[0]
	for index, num := range candles {
		if index > 0 && num == largest {
			result++
		}
		if index > 0 && num != largest {
			break
		}
	}
	return result
}

func main() {
	num := []int32{1, 2, 3, 4, 32, 32}
	fmt.Println(birthdayCakeCandles(num))
}
