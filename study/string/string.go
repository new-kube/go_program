package main

import (
	"fmt"
	"strings"
)

func main() {

	hello := "hello"
	value := strings.Split(hello, ",")
	fmt.Printf("%s split by ',' to %v\n", hello, value)

	name := "复习优化第1天-xxxxxxx"
	items := strings.Split(name, "-")

	items = strings.Split(items[0], "第")

	fmt.Printf("%v\n", items)

	num := strings.Split(items[1], "天")
	fmt.Printf("%v, count=%d\n", num, len(num))

	x, y, z := "一", "十一", "二十一"
	fmt.Printf("len(%s)=%d\n", x, len(x))
	fmt.Printf("len(%s)=%d\n", y, len(y))
	fmt.Printf("len(%s)=%d\n", z, len(z))
	fmt.Printf("len(%s)=%d\n", x, len([]rune(x)))
	fmt.Printf("len(%s)=%d\n", y, len([]rune(y)))
	fmt.Printf("len(%s)=%d\n", z, len([]rune(z)))

	a := "十一"
	sa := strings.Split(a, "十")
	fmt.Printf("%v, count=%d\n", sa, len(sa))

	xxx := " \t\n hello, 世界 \t\n\t"
	yyy := strings.Trim(xxx, " \t\n")
	fmt.Printf("'%s' trim result: '%s'\n", xxx, yyy)

	hello = "一二三四五六七八九十"
	fmt.Printf("hello=%s, hello[4]=%s\n", hello, hello[0:4])

	name = "一二三四五六七八九十"
	nameRune := []rune(name)
	fmt.Printf("string(nameRune[:4]) = %s\n", string(nameRune[:4]))
}
