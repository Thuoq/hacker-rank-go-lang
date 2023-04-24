package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	ch := make(chan int)
	//go func() {
	//	ch <- 1
	//	fmt.Println("Hello 2")
	//}()
	//fmt.Println(<-ch)
	go func() {
		for i := 1; i <= 10; i++ {
			// time.Sleep(time.Second * 1)
			ch <- i
		}
	}()
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}

}
