package main

import "fmt"

func main() {
	i, j := 29, 59
	p := &i

	i = 20

	fmt.Println(*p, i, j)
}
