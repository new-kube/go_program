package main

import "fmt"

// 字符串是不可修改的。
// s := "12345"
// s[2] = 'a'
// 根据这个特性，字符串的[]操作，只需要移动指针和长度就行，而不需要重新申请空间。
// x := "hello, world"
// y := x[5:] // ,world
// z := x[:5] // hello
// 都指向x的某个位置，然后长度不一样而已。

// 注意，``内部的字符串不会转义，是什么就是什么。
// x := `abc\n\n\r\n`
// 不做转义，所以多行字符串也可以，通常go的源码就是这样处理的。

// %[<n>] 占位符，其中<n>为后面的参数顺序，从1开始。
// %# 用于后面的 %d/%o/%x的添加前缀，注意不能和%d联合使用
// 只能是 %o/%x/%X

func modify(s string, i int, v byte) string {
	//s[i] = v // 编译失败，不能修改字符串内容
	return s
}

func testAddr() {
	x := "hello, world"
	y := x[5:]
	z := x[:5]
	a := x[:]
	b := x[:5] + "-world"

	fmt.Printf("x=%s, addr=%v\n", x, &x)
	fmt.Printf("y=%s, addr=%v\n", y, &y)
	fmt.Printf("z=%s, addr=%v\n", z, &z)
	fmt.Printf("a=%s, addr=%v\n", a, &a)
	fmt.Printf("b=%s, addr=%v\n", b, &b)
}

func testHat() {
	x := `abc\n\r\t\\%d%s`

	fmt.Println(x)
}

func testPlaceHolder() {
	x := 123
	y := "string"

	// 占位符为[<n>], 其中n位第几个参数。而#代表数字的八进制或者十六进制的前缀保留。
	fmt.Printf("x=%d, %[1]d, %[1]o, %#[1]o, %[1]x %#[1]x %#[1]X, y=%s\n", x, y)
}

func main() {
	testAddr()
	testHat()
	testPlaceHolder()
}
