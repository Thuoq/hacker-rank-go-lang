package main

import "fmt"

func staircase(n int) {
	// Write your code here
	for i := 0; i < n; i++ {
		for j := n; j >= 0; j-- {
			if i >= j {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func main() {
	n := 6
	staircase(n)
}
