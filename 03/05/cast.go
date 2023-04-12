package main

import (
	"fmt"
	"strconv"
)

// fmt.Sprintf()  strconv.Itoa()
// strconv.FormatInt() strconv.FormatUint()

// strconv.Atoi()
// strconv.ParseInt() strconv.ParseUint()
//

func test_printf() {
	x := 123
	y := fmt.Sprintf("%d", x)

	fmt.Printf("x = %d, y = %q\n", x, y)
}

func testStrconv() {
	// format.
	x := 123
	y := strconv.Itoa(x)
	fmt.Printf("x = %d, y = %q\n", x, y)

	x = 100
	y = strconv.FormatInt(int64(x), 10)
	fmt.Printf("x = %d, y = %q\n", x, y)

	y = fmt.Sprintf("%o", x)
	fmt.Printf("x = %d, y = %q\n", x, y)

	y = strconv.FormatInt(int64(x), 8)
	fmt.Printf("x = %d, y = %q\n", x, y)

	y = fmt.Sprintf("%x", x)
	fmt.Printf("x = %d, y = %q\n", x, y)

	y = strconv.FormatInt(int64(x), 16)
	fmt.Printf("x = %d, y = %q\n", x, y)

	// parse.
	a, err := strconv.ParseInt(y, 16, 32)
	if err != nil {
		fmt.Printf("error parsing, error: %v\n", err)
	}
	fmt.Printf("a = %d, y = %q\n", a, y)

	b := "100"
	c, err := strconv.Atoi(b)
	if err != nil {
		fmt.Printf("error parsing, error: %v\n", err)
	}
	fmt.Printf("c = %d, b = %q\n", c, b)
}

func main() {
	test_printf()
	testStrconv()
}
