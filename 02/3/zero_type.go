package main

import (
	"fmt"
)

type A struct{}

func main() {
	p := new(struct{})
	q := new(struct{})

	fmt.Print(p == q)
	// $ go version
	// go version go1.19.1 darwin/amd64
	// $ ./zero_type
	// false
}
