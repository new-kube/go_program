package main

import (
	"crypto/sha256"
	"fmt"
)

// array是固定长度的相同类型的序列。

func array() {
	var a [3]int // 初始化一个数组，元素都为0

	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

func array_init() {
	var a [3]int // 默认初始化为0
	var b [3]int = [3]int{1, 2, 3}
	var c [3]int = [3]int{1, 2} // 剩余用0填充

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	x := [...]int{1, 2, 3} // 注意省略号根据后面的值数量来确定
	fmt.Printf("x=%v, type=%T\n", x, x)

	// 数组的长度是数组类型的一部分，所以 不同长度的数组，是不同的数组类型。
	// 数组长度必须是常量表达式，即需要在编译的时候确定。

	// 初始化支持 索引: 值 的方式 如下：
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "E", GBP: "L", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
	// 上面的赋值，索引不需要按照顺序出现，并且中间可以缺失，缺失补充为零值，
	r := [...]int{99: -1}
	fmt.Printf("r, type=%T\n", r)
}

func array_compare() {
	// 数组具有相同类型才可比较，否则编译错误，当具有相同类型的时候:
	// == 所有元素都相同为true
	// != 至少有一个元素不同则为true

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)

	d := [3]int{1, 2}
	//fmt.Println(a == d) // 编译错误，不同类型的数组，无法比较
	fmt.Printf("d type=%T\n", d)
}

func crypto_compare() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func func_with_array(array [100]int) {
	// 数组作为参数是值传递。即意味着会存在拷贝，复制，并且低效，函数内外互不影响。
	// 这点和其他语言大不相同，如果考虑效率，可以采用传数组指针，这样就避免了复制。
}

func zero(array *[100]int) {
	for i := range array {
		array[i] = 0
	}
}

func zero2(array *[100]int) {
	*array = [100]int{}
}

func main() {
	crypto_compare()
}
