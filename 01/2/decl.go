package main

import "fmt"

// declare usage

func declare() {
	s := "" // 函数作用域中，不在包级别作用域中，非常常用。短小简洁。

	var ss string // 利用默认初始化给变量赋值为空字符串，可以用在函数或包作用域中。

	var sss = "" // 不太常见，通常在声明多个变量时使用

	var a, b, c = 1, "", 12.34

	var ssss string = "" // 通常用于，短变量声明的类型和想要声明类型不匹配的时候，这时是必须的。如下:

	k := 1           // int
	var kk int64 = 1 // int64, 这种情况就无法使用短变量声明了，因为默认数字是整型(int), 而不是int64.

	fmt.Println(s, ss, sss, a, b, c, ssss, k, kk)
}

func main() {
	declare()

	var x []int // 空slice，并不是nil，这个需要注意。

	fmt.Printf("x=%v\n", x)

	const y = 12

	// 常量无法获取地址，会出现编译错误。
	//z := &y

	type t struct{}

	//
	a := new(struct{})
	b := new(struct{})

	fmt.Printf("a == b? %t\n", a == b)

	//
	c := new(t)
	d := new(t)

	fmt.Printf("c == d? %t\n", c == d)

}
