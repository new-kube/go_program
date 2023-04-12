package main

import (
	"fmt"
	"unicode/utf8"
)

// utf编码。
// 兼容ascii码。
// go的源码用utf-8编码，go中的字符串优先采用utf-8编码。
// 如此，带有非一字节的字符，字符串的长度和字符串本身的逻辑长度并不一致。
// 	1. 通常可以使用这个函数获取真实的字符串长度：utf8.RuneCountInString(str))
//  上述方法按照utf-8编码规则获取的长度，是真实的长度。
// 当出现超过一字节的字符组成的字符串，可以使用如下的方式：
//  1. 原本的值，如中文：世界。
//  2. 字符串的码点，即对应的字符的utf-8的码点，形式如下：
//    \uhhhh -> 16位 4个16进制的字符码点
//    \Uhhhhhhhh -> 32位。8个16进制的字符码点
//  也可以用\xhh表示，十六进制的两个字符

func utf_encode() {
	hello := "世界"
	x := "\xe4\xb8\x96\xe7\x95\x8c"
	y := "\u4e16\u754c"
	z := "\U00004e16\U0000754c"

	fmt.Printf("hello=%s, len(hello)=%d, real len=%d\n", hello, len(hello), utf8.RuneCountInString(hello))
	fmt.Printf("x=%s, len(x)=%d, real len=%d\n", x, len(x), utf8.RuneCountInString(x))
	fmt.Printf("y=%s, len(y)=%d, real len=%d\n", y, len(y), utf8.RuneCountInString(y))
	fmt.Printf("z=%s, len(z)=%d, real len=%d\n", z, len(z), utf8.RuneCountInString(z))

	/* 输出
	hello=世界, len(hello)=6
	x=世界, len(x)=6
	y=世界, len(y)=6
	z=世界, len(z)=6
	*/
}

func example() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}

// 逐个打印字符串，可以采用
// 1. utf8.DecodeRuneInString()方法
// 2. for s := range str的方法。

func print_char(s string) {
	for i := 0; i < len(s); i++ {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, c := range s {
		fmt.Printf("%d\t%c\n", i, c)
	}
}

func main() {
	utf_encode()
	example()

	s := "hello, 世界"
	print_char(s)
}
