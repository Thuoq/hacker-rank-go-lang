package main

import "fmt"

type Point struct {
	X, Y int
}

var (
	p1 = Point{1, 2}
	p2 = Point{X: 1}
	p3 = &Point{4, 10}
)

func main() {
	//p1 := Point{1, 2}
	//fmt.Println(p1)
	//
	//p2 := &p1
	//
	//p2.X = 120
	fmt.Println(p1, p2, p3)
	p3.Y = 100

}
