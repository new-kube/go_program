package main

import (
	"fmt"
	"math"
)

func main() {
	i := 10
	fmt.Printf("i = %d, 0x%x, 0%o, %b\n", i, i, i, i)

	j := 1.233
	fmt.Printf("j = %f, %g, %e\n", j, j, j)

	k := true
	fmt.Printf("k = %t\n", k)

	a := 'c'
	fmt.Printf("a = %c\n", a)

	b := "hello, world"
	fmt.Printf("b = %s\n", b)

	c := "fuck"
	fmt.Printf("c = %q\n", c)

	type hello struct {
		A string
		b int
		C bool
		H *hello
	}

	h := hello{
		A: "hello",
		b: 123,
		C: true,
	}
	h.H = &h
	fmt.Printf("h = %v\n", h)

	fmt.Printf("h's type = %T\n", h)
	fmt.Printf("I am a %%\n")

	hx := &h
	fmt.Printf("&h = %v, %v, %+v\n", &h, hx, &h)

	// %f/%g/%e
	// %f的小数精度是 6 位, 可以使用 %.<n>f方式指定后面的位数。
	// %g的小数精度为 15位
	// %e为科学计数法，小数点后精度为6位和%f相同。

	x := math.Pi
	fmt.Printf("pi = %.8f, %g, %e\n", x, x, x)

}
