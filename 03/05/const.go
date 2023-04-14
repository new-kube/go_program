package main

import "fmt"

// 常量。

// 编译器就确定了值，
//  1. 可以提前计算，并优化。
// 	2. 可以作为数组的长度，作为声明，因为数组在编译期也要确定其值。

// 声明
//  1. 后面没有值，继承前面的声明的值。
//  2. iota 可以通过前一个的计算公式推导出后面的值，可以是比较复杂的公式，其中iota从0开始递增。

func testArray() {
	const IPv4Len = 4
	var p [IPv4Len]byte

	fmt.Printf("p=%v\n", p)
}

func testOmit() {
	const (
		a = 1
		b
		c = 2
		d
	)

	fmt.Printf("a=%d b=%d c=%d d=%d\n", a, b, c, d)

	const (
		x = "hello"
		y
		z = "world"
	)

	fmt.Printf("x=%s y=%s z=%s\n", x, y, z)
}

func testIotaSimple() {
	const (
		Sun = iota
		Mon
		Tue
		Wed
		Thu
		Fri
		Sat
	)

	fmt.Printf("sun=%d mon=%d tue=%d wed=%d thu=%d fri=%d sat=%d\n", Sun, Mon, Tue, Wed, Thu, Fri, Sat)
}

func testIotaComplex() {
	// 无类型常量，会根据赋值语句左边的类型来赋值对应的类型。从而能保证动态特性和精度。
	// 无类型常量，通常在const声明的时候不指定类型，就是无类型常量。通常精度会很高，能达到256位或者更高。
	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB // 超过了 1<<32
		PiB
		EiB
		ZiB // 超过了 1 << 64 超过了%d的最大值，无法打印。
		YiB
	)

	fmt.Printf("kib=%d, mib=%d, gib=%d, tib=%d, pib=%d, eib=%d, zib=%d, yib=%d\n", KiB, MiB, GiB, TiB, PiB, EiB, ZiB, YiB)
}

func main() {
	testArray()
	testOmit()
	testIotaSimple()
	testIotaComplex()
}
