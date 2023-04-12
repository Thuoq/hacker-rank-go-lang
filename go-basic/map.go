package main

import "fmt"

type point struct {
	Lat, Long float64
}

var mp map[string]point

func main() {
	mp = make(map[string]point)
	mp["hello"] = point{1, 2}
	fmt.Println(dict)

}
