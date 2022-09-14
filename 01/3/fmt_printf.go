package main

import "fmt"

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
	}

	h := hello{
		A: "hello",
		b: 123,
		C: true,
	}
	fmt.Printf("h = %v\n", h)

	fmt.Printf("h's type = %T\n", h)
	fmt.Printf("I am a %%\n")

}
