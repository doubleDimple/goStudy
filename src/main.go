package main

import (
	"fmt"
	"geometry/rectangle"
)

func main() {

	var width float64 = 24
	var length float64 = 35
	fmt.Print(width, length)
	area := rectangle.Area(width, length)
	fmt.Println("thw area is:", area)

}
