package main

import "fmt"

func main() {

	x := make([]int, 0)
	for i := 0; i < 10; i++ {
		x = append(x, i)
	}
	y := x
	fmt.Printf("x: %p, y: %p, %v\n", x, y, y)

	x = make([]int, 0)
	fmt.Printf("x: %p, y: %p, %v\n", x, y, y)
}
