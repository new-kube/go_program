package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	range_test()
}

func range_test() {
	a := []int{11, 12, 13, 14, 15}
	for i, v := range a { // 注意i的值，会从0开始递增。
		fmt.Printf("a[%d]=%d\n", i, v)
	}
}
