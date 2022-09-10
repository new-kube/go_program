package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		// 注意i++是语句，不是表达式，和c/c++不同，所以不能写成 j = i++; 的方式。主要是为了避免c/c++的这种迷惑行为。
		// 没有++i或者--i的形式，这在go中不合法。
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
