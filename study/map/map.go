package main

import "fmt"

func main() {
	hello := make(map[string]int)
	hello["hello"] = 1

	delete(hello, "world")

	fmt.Println("delete map not exist key.")
}
